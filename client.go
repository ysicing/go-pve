package pve

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	"golang.org/x/time/rate"
)

const (
	defaultBaseURL = "https://localhost:8006/"
	apiVersionPath = "api2/json/"
	userAgent      = "go-pve"

	// Default timeout and retry settings
	defaultTimeout   = 30 * time.Second
	defaultRetries   = 3
	defaultRateLimit = rate.Limit(10) // 10 requests per second
)

// AuthType represents authentication type
type AuthType int

const (
	PasswordAuth AuthType = iota
	TokenAuth
)

// RateLimiter interface for controlling request rate
type RateLimiter interface {
	Wait(ctx context.Context) error
}

// AuthOptions holds authentication configuration
type AuthOptions struct {
	Username     string
	Password     string
	TokenID      string
	TokenSecret  string
	AuthType     AuthType
	CSRFPreToken string
}

// Client represents a Proxmox VE API client
type Client struct {
	// HTTP client
	client *retryablehttp.Client

	// Base URL for API requests
	baseURL *url.URL

	// Authentication
	authOptions *AuthOptions
	authToken   string
	authCookie  string
	csrfToken   string

	// Rate limiting
	limiter RateLimiter

	// API services
	Cluster  *ClusterService
	Nodes    *NodesService
	VMs      *VMsService
	Storage  *StorageService
	Tasks    *TasksService
	Auth     *AuthService
	Version  *VersionService

	// User agent
	UserAgent string
}

// NewClient creates a new PVE API client
func NewClient(baseURL string, authOptions *AuthOptions, options ...ClientOptionFunc) (*Client, error) {
	// Parse base URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	if !u.IsAbs() {
		return nil, errors.New("base URL must be absolute")
	}
	if !strings.HasSuffix(u.Path, "/") {
		u.Path = u.Path + "/"
	}

	// Set default auth type
	if authOptions.AuthType == 0 {
		if authOptions.TokenID != "" && authOptions.TokenSecret != "" {
			authOptions.AuthType = TokenAuth
		} else {
			authOptions.AuthType = PasswordAuth
		}
	}

	// Create HTTP client
	httpClient := retryablehttp.NewClient()
	httpClient.HTTPClient.Timeout = defaultTimeout
	httpClient.Logger = nil // Disable default logging
	httpClient.RetryMax = defaultRetries
	httpClient.RetryWaitMin = 1 * time.Second
	httpClient.RetryWaitMax = 30 * time.Second

	// Create client
	c := &Client{
		client:      httpClient,
		baseURL:     u,
		authOptions: authOptions,
		limiter:     rate.NewLimiter(defaultRateLimit, 1),
		UserAgent:   userAgent,
	}

	// Apply options
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	// Initialize services
	c.Cluster = &ClusterService{client: c}
	c.Nodes = &NodesService{client: c}
	c.VMs = &VMsService{client: c}
	c.Storage = &StorageService{client: c}
	c.Tasks = &TasksService{client: c}
	c.Auth = &AuthService{client: c}
	c.Version = &VersionService{client: c}

	return c, nil
}

// ClientOptionFunc configures the client
type ClientOptionFunc func(*Client) error

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		c.client.HTTPClient = httpClient
		return nil
	}
}

// WithRateLimiter sets a custom rate limiter
func WithRateLimiter(limiter RateLimiter) ClientOptionFunc {
	return func(c *Client) error {
		c.limiter = limiter
		return nil
	}
}

// WithUserAgent sets a custom user agent
func WithUserAgent(ua string) ClientOptionFunc {
	return func(c *Client) error {
		c.UserAgent = ua
		return nil
	}
}

// NewRequest creates an HTTP request
func (c *Client) NewRequest(method, path string, opt interface{}, options ...RequestOptionFunc) (*http.Request, error) {
	u := c.baseURL.String() + apiVersionPath + path

	// Add query parameters
	if opt != nil {
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u += "?" + q.Encode()
	}

	// Create request
	req, err := http.NewRequest(method, u, nil)
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")

	// Apply request options
	for _, fn := range options {
		if err := fn(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}

// Do executes an HTTP request
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	// Rate limiting
	if c.limiter != nil {
		if err := c.limiter.Wait(req.Context()); err != nil {
			return nil, err
		}
	}

	// Ensure authentication
	if c.authCookie == "" && c.authToken == "" {
		if err := c.authenticate(); err != nil {
			return nil, err
		}
	}

	// Add authentication headers
	if c.authCookie != "" {
		req.Header.Set("Cookie", c.authCookie)
	}
	if c.authToken != "" {
		req.Header.Set("Authorization", c.authToken)
	}
	if c.csrfToken != "" {
		req.Header.Set("CSRFPreventionToken", c.csrfToken)
	}

	// Execute request using standard http client
	httpClient := c.client.HTTPClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Create response wrapper
	response := &Response{
		Response: resp,
	}

	// Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	response.Body = body

	// Handle authentication errors
	if resp.StatusCode == http.StatusUnauthorized {
		// Try to re-authenticate
		if c.authOptions.AuthType == PasswordAuth {
			c.authCookie = ""
			c.authToken = ""
			if err := c.authenticate(); err == nil {
				// Clone the request and retry
				newReq, err := http.NewRequest(req.Method, req.URL.String(), nil)
				if err == nil {
					newReq.Header = make(http.Header)
					for k, v := range req.Header {
						newReq.Header[k] = v
					}
					return c.Do(newReq, v)
				}
			}
		}
	}

	// Parse response
	if v != nil && len(body) > 0 && resp.StatusCode != http.StatusNoContent {
		if r, ok := v.(*[]byte); ok {
			*r = body
			return response, nil
		}

		// Try to unmarshal as JSON
		if contentType := resp.Header.Get("Content-Type"); contentType == "application/json" {
			decoder := json.NewDecoder(bytes.NewReader(body))
			if err := decoder.Decode(v); err != nil {
				return response, err
			}
		}
	}

	return response, nil
}

// authenticate handles authentication with Proxmox VE API
func (c *Client) authenticate() error {
	switch c.authOptions.AuthType {
	case PasswordAuth:
		return c.passwordAuth()
	case TokenAuth:
		return c.tokenAuth()
	default:
		return errors.New("unknown authentication type")
	}
}

// passwordAuth performs password authentication
func (c *Client) passwordAuth() error {
	reqData := map[string]string{
		"username": c.authOptions.Username,
		"password": c.authOptions.Password,
	}

	if c.authOptions.CSRFPreToken != "" {
		reqData["CSRFPreToken"] = c.authOptions.CSRFPreToken
	}

	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.baseURL.String()+apiVersionPath+"access/ticket", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req, nil)
	if err != nil {
		return err
	}

	// Parse response
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return err
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		return errors.New("invalid authentication response")
	}

	// Extract tokens
	if ticket, ok := data["ticket"].(string); ok {
		c.authCookie = "PVEAuthCookie=" + ticket
	}
	if csrf, ok := data["CSRFPreventionToken"].(string); ok {
		c.csrfToken = csrf
	}

	return nil
}

// tokenAuth performs token-based authentication
func (c *Client) tokenAuth() error {
	// For token auth, we don't need to authenticate via API
	// Just construct the authorization header
	username := c.authOptions.Username
	if username == "" {
		return errors.New("username is required for token authentication")
	}

	c.authToken = "PVEAPIToken=" + username + "@" + c.authOptions.TokenID + "!" + c.authOptions.TokenSecret

	return nil
}

// Response wraps http.Response and provides access to response data
type Response struct {
	*http.Response

	Body []byte
}

// String returns response body as string
func (r *Response) String() string {
	return string(r.Body)
}

// ParseError parses error from API response
func (c *Client) ParseError(r *Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return nil
	}

	var errResult struct {
		Errors []string `json:"errors"`
		Data   string   `json:"data"`
	}

	if err := json.Unmarshal(r.Body, &errResult); err != nil {
		return fmt.Errorf("API error (status %d): %s", r.StatusCode, r.String())
	}

	if len(errResult.Errors) > 0 {
		return errors.New(strings.Join(errResult.Errors, ", "))
	}

	if errResult.Data != "" {
		return errors.New(errResult.Data)
	}

	return fmt.Errorf("API error (status %d)", r.StatusCode)
}

// parseID converts various ID types to string
func parseID(id interface{}) (string, error) {
	switch v := id.(type) {
	case int:
		return strconv.Itoa(v), nil
	case string:
		return v, nil
	default:
		return "", errors.New("invalid ID type")
	}
}

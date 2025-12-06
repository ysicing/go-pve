package pve

import "fmt"

// AuthService handles authentication-related API operations
type AuthService struct {
	client *Client
}

// Login performs login authentication
func (s *AuthService) Login(username, password string) (*Ticket, error) {
	reqData := map[string]interface{}{
		"username": username,
		"password": password,
	}

	req, err := s.client.NewRequest("POST", "access/ticket", reqData)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *Ticket
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// Logout performs logout
func (s *AuthService) Logout() error {
	req, err := s.client.NewRequest("POST", "access/logout", nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// GetTicketInfo retrieves current ticket information
func (s *AuthService) GetTicketInfo() (*Ticket, error) {
	req, err := s.client.NewRequest("GET", "access/ticket", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *Ticket
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetPermissions retrieves user permissions
func (s *AuthService) GetPermissions(path string) (map[string]interface{}, error) {
	req, err := s.client.NewRequest("GET", "access/permissions", map[string]interface{}{
		"path": path,
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		Data map[string]interface{}
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetUsers retrieves all users
func (s *AuthService) GetUsers() ([]map[string]interface{}, error) {
	req, err := s.client.NewRequest("GET", "access/users", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []map[string]interface{}
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetUser retrieves a specific user
func (s *AuthService) GetUser(userid string) (map[string]interface{}, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("access/users/%s", userid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data map[string]interface{}
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// CreateUser creates a new user
func (s *AuthService) CreateUser(userid, password, email string) error {
	req, err := s.client.NewRequest("POST", "access/users", map[string]interface{}{
		"userid":  userid,
		"password": password,
		"email":   email,
	})
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// UpdateUser updates user information
func (s *AuthService) UpdateUser(userid, email string) error {
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("access/users/%s", userid), map[string]interface{}{
		"email": email,
	})
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// DeleteUser deletes a user
func (s *AuthService) DeleteUser(userid string) error {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("access/users/%s", userid), nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// GetRoles retrieves all roles
func (s *AuthService) GetRoles() ([]map[string]interface{}, error) {
	req, err := s.client.NewRequest("GET", "access/roles", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []map[string]interface{}
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

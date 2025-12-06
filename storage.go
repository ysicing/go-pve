package pve

import (
	"fmt"
)

// StorageService handles storage-related API operations
type StorageService struct {
	client *Client
}

// List retrieves all storage entities
func (s *StorageService) List(options *StorageListOptions) ([]*Storage, error) {
	req, err := s.client.NewRequest("GET", "storage", options)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*Storage
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// Get retrieves a specific storage by name
func (s *StorageService) Get(name string) (*Storage, error) {
	storages, err := s.List(nil)
	if err != nil {
		return nil, err
	}

	for _, storage := range storages {
		if storage.Storage == name {
			return storage, nil
		}
	}

	return nil, fmt.Errorf("storage %s not found", name)
}

// GetContent retrieves storage content
func (s *StorageService) GetContent(storageName string) ([]map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("storage/%s/content", storageName), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []map[string]any
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetContentByType retrieves storage content filtered by type
func (s *StorageService) GetContentByType(storageName, contentType string) ([]map[string]any, error) {
	content, err := s.GetContent(storageName)
	if err != nil {
		return nil, err
	}

	var filtered []map[string]any
	for _, item := range content {
		if ct, ok := item["content"].(string); ok && ct == contentType {
			filtered = append(filtered, item)
		}
	}

	return filtered, nil
}

// ListContent retrieves storage content with options
func (s *StorageService) ListContent(storageName string, options *StorageListOptions) ([]map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("storage/%s/content", storageName), options)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []map[string]any
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// Upload uploads a file to storage
func (s *StorageService) Upload(storageName, filename string, content []byte) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("storage/%s/content", storageName), map[string]any{
		"filename": filename,
		"content":  content,
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *Task
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// Download downloads a file from storage
func (s *StorageService) Download(storageName, volume string) ([]byte, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("storage/%s/content/%s", storageName, volume), nil)
	if err != nil {
		return nil, err
	}

	var result []byte
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteContent deletes content from storage
func (s *StorageService) DeleteContent(storageName, volume string) (*Task, error) {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("storage/%s/content/%s", storageName, volume), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *Task
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetDir retrieves directory listing
func (s *StorageService) GetDir(storageName string) ([]map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("storage/%s/dir", storageName), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []map[string]any
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetRRD retrieves storage RRD data
func (s *StorageService) GetRRD(storageName, timeframe string) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("storage/%s/rrddata", storageName), map[string]any{
		"timeframe": timeframe,
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		Data map[string]any
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

package pve

// VersionService handles version-related API operations
type VersionService struct {
	client *Client
}

// Get retrieves version information
func (s *VersionService) Get() (*Version, error) {
	req, err := s.client.NewRequest("GET", "version", nil)
	if err != nil {
		return nil, err
	}

	version := &Version{}
	_, err = s.client.Do(req, version)
	if err != nil {
		return nil, err
	}

	return version, nil
}

// GetAPT retrieves APT version information
func (s *VersionService) GetAPT() (map[string]interface{}, error) {
	req, err := s.client.NewRequest("GET", "apt/update", nil)
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

// GetPackages retrieves available packages
func (s *VersionService) GetPackages() ([]map[string]interface{}, error) {
	req, err := s.client.NewRequest("GET", "apt/versions", nil)
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

// Changelog retrieves changelog information for a package
func (s *VersionService) Changelog(packageName string) (string, error) {
	req, err := s.client.NewRequest("GET", "apt/changelog", map[string]interface{}{
		"package": packageName,
	})
	if err != nil {
		return "", err
	}

	var result string
	_, err = s.client.Do(req, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}

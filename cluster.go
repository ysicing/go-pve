package pve

// ClusterService handles cluster-related API operations
type ClusterService struct {
	client *Client
}

// Get retrieves cluster information
func (s *ClusterService) Get() (*Cluster, error) {
	req, err := s.client.NewRequest("GET", "cluster", nil)
	if err != nil {
		return nil, err
	}

	cluster := &Cluster{}
	_, err = s.client.Do(req, cluster)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

// Resources retrieves all cluster resources
func (s *ClusterService) Resources() ([]*ClusterResource, error) {
	req, err := s.client.NewRequest("GET", "cluster/resources", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*ClusterResource
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// ResourcesByType retrieves cluster resources filtered by type
func (s *ClusterService) ResourcesByType(resourceType string) ([]*ClusterResource, error) {
	resources, err := s.Resources()
	if err != nil {
		return nil, err
	}

	var filtered []*ClusterResource
	for _, r := range resources {
		if r.Type == resourceType {
			filtered = append(filtered, r)
		}
	}

	return filtered, nil
}

// GetResource retrieves a specific cluster resource
func (s *ClusterService) GetResource(resourceID string) (*ClusterResource, error) {
	req, err := s.client.NewRequest("GET", "cluster/resources/"+resourceID, nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *ClusterResource
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// Nodes returns cluster nodes information
func (s *ClusterService) Nodes() ([]*Node, error) {
	cluster, err := s.Get()
	if err != nil {
		return nil, err
	}

	return cluster.Nodes, nil
}

// Tasks retrieves cluster tasks
func (s *ClusterService) Tasks() ([]*Task, error) {
	req, err := s.client.NewRequest("GET", "cluster/tasks", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*Task
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// Status returns detailed cluster status
func (s *ClusterService) Status() (*Cluster, error) {
	cluster, err := s.Get()
	if err != nil {
		return nil, err
	}

	// Get additional status information
	resources, err := s.Resources()
	if err != nil {
		return cluster, nil
	}

	// Add resources to cluster for convenience
	for _, r := range resources {
		if r.Type == "node" {
			cluster.Nodes = append(cluster.Nodes, &Node{
				Name:     r.Name,
				Status:   r.Status,
				CPU:      r.CPU,
				Mem:      r.Mem,
				MaxMem:   r.MaxMem,
				Uptime:   r.Uptime,
			})
		}
	}

	return cluster, nil
}

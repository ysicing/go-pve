package pve

import (
	"fmt"
)

// NodesService handles node-related API operations
type NodesService struct {
	client *Client
}

// List retrieves all nodes
func (s *NodesService) List() ([]*Node, error) {
	req, err := s.client.NewRequest("GET", "nodes", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*Node
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// Get retrieves a specific node by name
func (s *NodesService) Get(name string) (*Node, error) {
	nodes, err := s.List()
	if err != nil {
		return nil, err
	}

	for _, node := range nodes {
		if node.Name == name {
			return node, nil
		}
	}

	return nil, fmt.Errorf("node %s not found", name)
}

// GetDetailed retrieves detailed node information
func (s *NodesService) GetDetailed(name string) (*NodeInfo, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s", name), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *NodeInfo
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetStatus retrieves node status
func (s *NodesService) GetStatus(name string) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/status", name), nil)
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

// GetVersion retrieves node version information
func (s *NodesService) GetVersion(name string) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/status/version", name), nil)
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

// GetConfig retrieves node configuration information
func (s *NodesService) GetConfig(name string) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/config", name), nil)
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

// CreateVNCShell creates a VNC shell for node access
func (s *NodesService) CreateVNCShell(name string) (map[string]any, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/vncshell", name), nil)
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

// GetSubscription retrieves node subscription information
func (s *NodesService) GetSubscription(name string) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/subscription", name), nil)
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

// GetSyslog retrieves node system log
func (s *NodesService) GetSyslog(name string, lines int) ([]string, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/syslog", name), map[string]any{
		"lines": lines,
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []string
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetRRD retrieves node RRD (Round Robin Database) data
func (s *NodesService) GetRRD(name, timeframe string) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/rrddata", name), map[string]any{
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

// GetTasks retrieves node tasks
func (s *NodesService) GetTasks(name string) ([]*Task, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/tasks", name), nil)
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

// StartNode starts a node (used for start/stop of services, not shutdown)
func (s *NodesService) Start(name string) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/status/start", name), nil)
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

// StopNode stops a node
func (s *NodesService) Stop(name string) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/status/stop", name), nil)
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

// ShutdownNode shuts down a node
func (s *NodesService) Shutdown(name string) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/status/shutdown", name), nil)
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

// RebootNode reboots a node
func (s *NodesService) Reboot(name string) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/status/reboot", name), nil)
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

// GetStorage retrieves storage information for a node
func (s *NodesService) GetStorage(name string) ([]*Storage, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/storage", name), nil)
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

// GetVMs retrieves all VMs on a node
func (s *NodesService) GetVMs(name string) ([]*VM, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc", name), nil)
	if err != nil {
		return nil, err
	}

	var resultLXC struct {
		Data []*VM
	}
	_, err = s.client.Do(req, &resultLXC)
	if err != nil {
		return nil, err
	}

	req, err = s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu", name), nil)
	if err != nil {
		return nil, err
	}

	var resultQEMU struct {
		Data []*VM
	}
	_, err = s.client.Do(req, &resultQEMU)
	if err != nil {
		return nil, err
	}

	// Combine results
	vms := append(resultLXC.Data, resultQEMU.Data...)
	return vms, nil
}

// GetNetstat retrieves network connection statistics for a node
func (s *NodesService) GetNetstat(name string) ([]map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/netstat", name), nil)
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

// GetQEMUVMs retrieves all QEMU VMs on a node
func (s *NodesService) GetQEMUVMs(name string) ([]*VM, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu", name), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*VM
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetLXCContainers retrieves all LXC containers on a node
func (s *NodesService) GetLXCContainers(name string) ([]*VM, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc", name), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*VM
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// CreateVZDumpBackup creates a backup task using vzdump
func (s *NodesService) CreateVZDumpBackup(name string, options *VZDumpOptions) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/vzdump", name), options)
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

// ExtractVZDumpConfig extracts vzdump backup configuration
func (s *NodesService) ExtractVZDumpConfig(name, volume string) (string, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/vzdump/extractconfig", name), map[string]any{
		"volume": volume,
	})
	if err != nil {
		return "", err
	}

	var result struct {
		Data string
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return "", err
	}

	return result.Data, nil
}

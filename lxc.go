package pve

import (
	"fmt"
)

// LXCService handles LXC-specific container operations
type LXCService struct {
	client *Client
}

// List retrieves all LXC containers on a node
func (s *LXCService) List(node string) ([]*VM, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc", node), nil)
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

// Get retrieves a specific LXC container
func (s *LXCService) Get(node string, vmid int) (*VMStatus, error) {
	status, err := s.GetStatus(node, vmid)
	if err != nil {
		return nil, err
	}

	return status, nil
}

// GetStatus retrieves LXC container current status
func (s *LXCService) GetStatus(node string, vmid int) (*VMStatus, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc/%d/status/current", node, vmid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *VMStatus
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetConfig retrieves LXC container configuration
func (s *LXCService) GetConfig(node string, vmid int) (*VMConfig, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc/%d/config", node, vmid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *VMConfig
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// UpdateConfig updates LXC container configuration
func (s *LXCService) UpdateConfig(node string, vmid int, config map[string]string) (*Task, error) {
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("nodes/%s/lxc/%d/config", node, vmid), config)
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

// Start starts an LXC container
func (s *LXCService) Start(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/status/start", node, vmid), nil)
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

// Stop stops an LXC container
func (s *LXCService) Stop(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/status/stop", node, vmid), nil)
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

// Shutdown gracefully shuts down an LXC container
func (s *LXCService) Shutdown(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/status/shutdown", node, vmid), nil)
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

// Reboot reboots an LXC container
func (s *LXCService) Reboot(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/status/reboot", node, vmid), nil)
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

// Suspend suspends an LXC container
func (s *LXCService) Suspend(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/status/suspend", node, vmid), nil)
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

// Resume resumes a suspended LXC container
func (s *LXCService) Resume(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/status/resume", node, vmid), nil)
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

// Delete deletes an LXC container
func (s *LXCService) Delete(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("nodes/%s/lxc/%d", node, vmid), nil)
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

// Migrate migrates an LXC container to another node
func (s *LXCService) Migrate(node string, vmid int, target string, options *MigrateOptions) (*Task, error) {
	params := map[string]any{
		"target": target,
	}

	if options != nil {
		if options.Online {
			params["online"] = 1
		}
		if options.Force {
			params["force"] = 1
		}
		if options.BWLimit > 0 {
			params["bwlimit"] = options.BWLimit
		}
		if options.TargetStorage != "" {
			params["targetstorage"] = options.TargetStorage
		}
		if options.Delete {
			params["delete"] = 1
		}
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/migrate", node, vmid), params)
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

// Clone clones an LXC container
func (s *LXCService) Clone(node string, vmid int, newID int, hostname string, full bool) (*Task, error) {
	params := map[string]any{
		"newid":    newID,
		"hostname": hostname,
	}

	if full {
		params["full"] = 1
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/clone", node, vmid), params)
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

// ResizeDisk resizes an LXC container disk
func (s *LXCService) ResizeDisk(node string, vmid int, disk string, size string) (*Task, error) {
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("nodes/%s/lxc/%d/resize", node, vmid), map[string]string{
		"disk": disk,
		"size": size,
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

// ListSnapshots lists LXC container snapshots
func (s *LXCService) ListSnapshots(node string, vmid int) ([]*VMSnapshot, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc/%d/snapshot", node, vmid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []*VMSnapshot
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// CreateSnapshot creates an LXC container snapshot
func (s *LXCService) CreateSnapshot(node string, vmid int, name, description string) (*Task, error) {
	params := map[string]any{
		"snapname":    name,
		"description": description,
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/snapshot", node, vmid), params)
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

// DeleteSnapshot deletes an LXC container snapshot
func (s *LXCService) DeleteSnapshot(node string, vmid int, snapshotName string) (*Task, error) {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("nodes/%s/lxc/%d/snapshot/%s", node, vmid, snapshotName), nil)
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

// RollbackSnapshot rolls back to an LXC container snapshot
func (s *LXCService) RollbackSnapshot(node string, vmid int, snapshotName string) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/snapshot/%s/rollback", node, vmid, snapshotName), nil)
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

// GetVNCProxy gets VNC proxy information for an LXC container
func (s *LXCService) GetVNCProxy(node string, vmid int, websocket bool) (map[string]any, error) {
	params := map[string]any{}
	if websocket {
		params["websocket"] = 1
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/vncproxy", node, vmid), params)
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

// GetInterfaces retrieves LXC container network interfaces
func (s *LXCService) GetInterfaces(node string, vmid int) ([]NetworkInterface, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc/%d/interfaces", node, vmid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []NetworkInterface
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// EnterContainer enters an LXC container (creates a shell session)
func (s *LXCService) EnterContainer(node string, vmid int) (map[string]any, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/lxc/%d/termproxy", node, vmid), nil)
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

// GetPending retrieves pending LXC container configuration changes
func (s *LXCService) GetPending(node string, vmid int) (map[string]any, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc/%d/pending", node, vmid), nil)
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

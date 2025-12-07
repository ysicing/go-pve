package pve

import (
	"fmt"
)

// QEMUService handles QEMU-specific VM operations
type QEMUService struct {
	client *Client
}

// List retrieves all QEMU VMs across all nodes
func (s *QEMUService) List(node string) ([]*VM, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu", node), nil)
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

// Get retrieves a specific QEMU VM
func (s *QEMUService) Get(node string, vmid int) (*VMStatus, error) {
	status, err := s.GetStatus(node, vmid)
	if err != nil {
		return nil, err
	}

	return status, nil
}

// GetStatus retrieves QEMU VM current status
func (s *QEMUService) GetStatus(node string, vmid int) (*VMStatus, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/status/current", node, vmid), nil)
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

// GetConfig retrieves QEMU VM configuration
func (s *QEMUService) GetConfig(node string, vmid int) (*VMConfig, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/config", node, vmid), nil)
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

// UpdateConfig updates QEMU VM configuration
func (s *QEMUService) UpdateConfig(node string, vmid int, config map[string]string) (*Task, error) {
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("nodes/%s/qemu/%d/config", node, vmid), config)
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

// Start starts a QEMU VM
func (s *QEMUService) Start(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/status/start", node, vmid), nil)
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

// Stop stops a QEMU VM
func (s *QEMUService) Stop(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/status/stop", node, vmid), nil)
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

// Shutdown gracefully shuts down a QEMU VM
func (s *QEMUService) Shutdown(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/status/shutdown", node, vmid), nil)
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

// Reboot reboots a QEMU VM
func (s *QEMUService) Reboot(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/status/reboot", node, vmid), nil)
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

// Reset hard resets a QEMU VM
func (s *QEMUService) Reset(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/status/reset", node, vmid), nil)
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

// Suspend suspends a QEMU VM
func (s *QEMUService) Suspend(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/status/suspend", node, vmid), nil)
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

// Resume resumes a suspended QEMU VM
func (s *QEMUService) Resume(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/status/resume", node, vmid), nil)
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

// Delete deletes a QEMU VM
func (s *QEMUService) Delete(node string, vmid int) (*Task, error) {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("nodes/%s/qemu/%d", node, vmid), nil)
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

// Migrate migrates a QEMU VM to another node
func (s *QEMUService) Migrate(node string, vmid int, target string, options *MigrateOptions) (*Task, error) {
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
		if options.MigrationNetwork != "" {
			params["migration_network"] = options.MigrationNetwork
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

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/migrate", node, vmid), params)
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

// Clone clones a QEMU VM
func (s *QEMUService) Clone(node string, vmid int, newID int, name string, full bool) (*Task, error) {
	params := map[string]any{
		"newid": newID,
		"name":  name,
	}

	if full {
		params["full"] = 1
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/clone", node, vmid), params)
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

// ResizeDisk resizes a QEMU VM disk
func (s *QEMUService) ResizeDisk(node string, vmid int, disk string, size string) (*Task, error) {
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("nodes/%s/qemu/%d/resize", node, vmid), map[string]string{
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

// ListSnapshots lists QEMU VM snapshots
func (s *QEMUService) ListSnapshots(node string, vmid int) ([]*VMSnapshot, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/snapshot", node, vmid), nil)
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

// CreateSnapshot creates a QEMU VM snapshot
func (s *QEMUService) CreateSnapshot(node string, vmid int, name, description string, vmstate bool) (*Task, error) {
	params := map[string]any{
		"snapname":    name,
		"description": description,
	}

	if vmstate {
		params["vmstate"] = 1
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/snapshot", node, vmid), params)
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

// DeleteSnapshot deletes a QEMU VM snapshot
func (s *QEMUService) DeleteSnapshot(node string, vmid int, snapshotName string) (*Task, error) {
	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("nodes/%s/qemu/%d/snapshot/%s", node, vmid, snapshotName), nil)
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

// RollbackSnapshot rolls back to a QEMU VM snapshot
func (s *QEMUService) RollbackSnapshot(node string, vmid int, snapshotName string) (*Task, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/snapshot/%s/rollback", node, vmid, snapshotName), nil)
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

// GetVNCProxy gets VNC proxy information for a QEMU VM
func (s *QEMUService) GetVNCProxy(node string, vmid int, websocket bool) (map[string]any, error) {
	params := map[string]any{}
	if websocket {
		params["websocket"] = 1
		params["generate-password"] = 1
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/vncproxy", node, vmid), params)
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

// SendMonitorCommand sends a command to QEMU monitor
func (s *QEMUService) SendMonitorCommand(node string, vmid int, command string) (string, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/monitor", node, vmid), map[string]string{
		"command": command,
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

// GetAgentInfo retrieves QEMU guest agent information
func (s *QEMUService) GetAgentInfo(node string, vmid int) (*GuestAgent, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/agent/info", node, vmid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *GuestAgent
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetAgentNetworkInterfaces retrieves network interfaces via QEMU guest agent
func (s *QEMUService) GetAgentNetworkInterfaces(node string, vmid int) ([]NetworkInterface, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/agent/network-get-interfaces", node, vmid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			Result []NetworkInterface `json:"result"`
		} `json:"data"`
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.Result, nil
}

// GetAgentFilesystemInfo retrieves filesystem information via QEMU guest agent
func (s *QEMUService) GetAgentFilesystemInfo(node string, vmid int) ([]FilesystemInfo, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/agent/get-fsinfo", node, vmid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			Result []FilesystemInfo `json:"result"`
		} `json:"data"`
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data.Result, nil
}

// ExecuteAgentCommand executes a command via QEMU guest agent
func (s *QEMUService) ExecuteAgentCommand(node string, vmid int, command []string) (*GuestExec, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/agent/exec", node, vmid), map[string]any{
		"command": command,
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *GuestExec
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetAgentExecStatus retrieves execution status from QEMU guest agent
func (s *QEMUService) GetAgentExecStatus(node string, vmid int, pid int) (*GuestExecResult, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/agent/exec-status", node, vmid), map[string]any{
		"pid": pid,
	})
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *GuestExecResult
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

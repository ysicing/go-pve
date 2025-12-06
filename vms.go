package pve

import (
	"fmt"
)

// VMsService handles VM-related API operations
type VMsService struct {
	client *Client
}

// List retrieves all VMs
func (s *VMsService) List(options *VMListOptions) ([]*VM, error) {
	req, err := s.client.NewRequest("GET", "cluster/resources", options)
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

	// Filter only VMs/LXCs
	var vms []*VM
	for _, vm := range result.Data {
		if vm.Type == "qemu" || vm.Type == "lxc" {
			vms = append(vms, vm)
		}
	}

	return vms, nil
}

// Get retrieves a specific VM by ID
func (s *VMsService) Get(vmid int) (*VM, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	// Get detailed status
	status, err := s.GetStatus(vmid)
	if err != nil {
		return vm, nil
	}

	vm.CPU = status.CPU
	vm.Mem = status.Mem
	vm.MaxMem = status.MaxMem
	vm.Disk = status.Disk
	vm.MaxDisk = status.MaxDisk
	vm.DiskRead = status.DiskRead
	vm.DiskWrite = status.DiskWrite
	vm.NetIn = status.NetIn
	vm.NetOut = status.NetOut
	vm.Uptime = status.Uptime
	vm.Status = status.Status

	return vm, nil
}

// GetVMResource retrieves VM resource information
func (s *VMsService) GetVMResource(vmid int) (*VM, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("cluster/resources/vm/%d", vmid), nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data *VM
	}
	_, err = s.client.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetStatus retrieves VM status information
func (s *VMsService) GetStatus(vmid int) (*VMStatus, error) {
	// First get the VM resource to find the node
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/%s/%d/status/current", vm.Node, vm.Type, vmid), nil)
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

// StartVM starts a VM
func (s *VMsService) Start(vmid int) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/status/start", vm.Node, vm.Type, vmid), nil)
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

// StopVM stops a VM
func (s *VMsService) Stop(vmid int) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/status/stop", vm.Node, vm.Type, vmid), nil)
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

// ShutdownVM shuts down a VM gracefully
func (s *VMsService) Shutdown(vmid int) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/status/shutdown", vm.Node, vm.Type, vmid), nil)
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

// RebootVM reboots a VM
func (s *VMsService) Reboot(vmid int) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/status/reboot", vm.Node, vm.Type, vmid), nil)
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

// SuspendVM suspends a VM
func (s *VMsService) Suspend(vmid int) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/status/suspend", vm.Node, vm.Type, vmid), nil)
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

// ResumeVM resumes a suspended VM
func (s *VMsService) Resume(vmid int) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/status/resume", vm.Node, vm.Type, vmid), nil)
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

// DeleteVM removes a VM
func (s *VMsService) Delete(vmid int) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("nodes/%s/%s/%d", vm.Node, vm.Type, vmid), nil)
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

// GetConfig retrieves VM configuration
func (s *VMsService) GetConfig(vmid int) (*VMConfig, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/%s/%d/config", vm.Node, vm.Type, vmid), nil)
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

// UpdateConfig updates VM configuration
func (s *VMsService) UpdateConfig(vmid int, config map[string]string) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	// Convert config map to URL values
	values := make(map[string]string)
	for k, v := range config {
		values[k] = v
	}

	req, err := s.client.NewRequest("PUT", fmt.Sprintf("nodes/%s/%s/%d/config", vm.Node, vm.Type, vmid), values)
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

// ListSnapshots lists VM snapshots
func (s *VMsService) ListSnapshots(vmid int) ([]*VMSnapshot, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/%s/%d/snapshot", vm.Node, vm.Type, vmid), nil)
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

// CreateSnapshot creates a VM snapshot
func (s *VMsService) CreateSnapshot(vmid int, name, description string) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/snapshot", vm.Node, vm.Type, vmid), map[string]string{
		"snapname":    name,
		"description": description,
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

// DeleteSnapshot deletes a VM snapshot
func (s *VMsService) DeleteSnapshot(vmid int, snapshotName string) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("DELETE", fmt.Sprintf("nodes/%s/%s/%d/snapshot/%s", vm.Node, vm.Type, vmid, snapshotName), nil)
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

// RollbackSnapshot rolls back to a VM snapshot
func (s *VMsService) RollbackSnapshot(vmid int, snapshotName string) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/snapshot/%s/rollback", vm.Node, vm.Type, vmid, snapshotName), nil)
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

// CloneVM clones a VM
func (s *VMsService) Clone(vmid int, newID int, name string) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/clone", vm.Node, vm.Type, vmid), map[string]any{
		"vmid":   newID,
		"name":   name,
		"full":   1,
		"target": vm.Node,
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

// GetVNCInfo retrieves VNC console information
func (s *VMsService) GetVNCInfo(vmid int) (map[string]any, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/%s/%d/vncproxy", vm.Node, vm.Type, vmid), nil)
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

// GetGuestAgentInfo retrieves guest agent information
func (s *VMsService) GetGuestAgentInfo(vmid int) (*GuestAgent, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/agent/get-guest-info", vm.Node, vmid), nil)
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

// ExecGuestCommand executes a command in the guest
func (s *VMsService) ExecGuestCommand(vmid int, command string) (*GuestExec, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/agent/exec", vm.Node, vmid), map[string]string{
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

// GetExecOutput retrieves output from a guest command execution
func (s *VMsService) GetExecOutput(vmid int, pid int) (*GuestExecResult, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/agent/exec/%d", vm.Node, vmid, pid), nil)
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

// Reset hard resets a QEMU VM
func (s *VMsService) Reset(vmid int) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	if vm.Type != "qemu" {
		return nil, fmt.Errorf("reset is only supported for QEMU VMs")
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/qemu/%d/status/reset", vm.Node, vmid), nil)
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

// ResizeDisk resizes a VM disk
func (s *VMsService) ResizeDisk(vmid int, disk string, size string) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("PUT", fmt.Sprintf("nodes/%s/%s/%d/resize", vm.Node, vm.Type, vmid), map[string]string{
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

// Migrate migrates a VM to another node
func (s *VMsService) Migrate(vmid int, target string, options *MigrateOptions) (*Task, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

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

	req, err := s.client.NewRequest("POST", fmt.Sprintf("nodes/%s/%s/%d/migrate", vm.Node, vm.Type, vmid), params)
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

// GetNetworkInterfaces retrieves VM network interfaces (QEMU with guest agent)
func (s *VMsService) GetNetworkInterfaces(vmid int) ([]NetworkInterface, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	if vm.Type == "qemu" {
		req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/agent/network-get-interfaces", vm.Node, vmid), nil)
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
	} else if vm.Type == "lxc" {
		req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/lxc/%d/interfaces", vm.Node, vmid), nil)
		if err != nil {
			return nil, err
		}

		var result struct {
			Data []NetworkInterface `json:"data"`
		}
		_, err = s.client.Do(req, &result)
		if err != nil {
			return nil, err
		}

		return result.Data, nil
	}

	return nil, fmt.Errorf("unsupported VM type: %s", vm.Type)
}

// GetFilesystemInfo retrieves VM filesystem information (QEMU with guest agent)
func (s *VMsService) GetFilesystemInfo(vmid int) ([]FilesystemInfo, error) {
	vm, err := s.GetVMResource(vmid)
	if err != nil {
		return nil, err
	}

	if vm.Type != "qemu" {
		return nil, fmt.Errorf("filesystem info is only supported for QEMU VMs with guest agent")
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("nodes/%s/qemu/%d/agent/get-fsinfo", vm.Node, vmid), nil)
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

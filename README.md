# go-pve

A comprehensive Go client library for the Proxmox Virtual Environment (PVE) API.

[![Go Reference](https://pkg.go.dev/badge/github.com/ysicing/go-pve.svg)](https://pkg.go.dev/github.com/ysicing/go-pve)
[![Go Report Card](https://goreportcard.com/badge/github.com/ysicing/go-pve)](https://goreportcard.com/report/github.com/ysicing/go-pve)

## Features

- ✅ **Full API Coverage** - Support for all major Proxmox VE API endpoints
- ✅ **Type-Safe Architecture** - Separate services for QEMU VMs and LXC containers
- ✅ **Dual Authentication** - Password and API Token authentication
- ✅ **Automatic Retry** - Built-in retry logic with exponential backoff
- ✅ **Rate Limiting** - Configurable rate limiting support
- ✅ **Comprehensive Types** - Rich type definitions for all API responses
- ✅ **Clean API** - Idiomatic Go interfaces
- ✅ **Production Ready** - Used in production environments

## Installation

```bash
go get github.com/ysicing/go-pve
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    pve "github.com/ysicing/go-pve"
)

func main() {
    // Create client with password authentication
    client, err := pve.NewClient(
        "https://pve.example.com:8006",
        &pve.AuthOptions{
            Username: "root@pam",
            Password: "your-password",
            AuthType: pve.PasswordAuth,
        },
    )
    if err != nil {
        log.Fatal(err)
    }

    // List all QEMU VMs on a node
    vms, err := client.QEMU.List("pve-node1")
    if err != nil {
        log.Fatal(err)
    }

    for _, vm := range vms {
        fmt.Printf("VM %d: %s (%s)\n", vm.ID, vm.Name, vm.Status)
    }
}
```

## Architecture

The library is organized into specialized services for different resource types:

```
Client
├── Cluster   - Cluster-wide operations
├── Nodes     - Node management
├── VMs       - Generic VM operations (all types)
├── QEMU      - QEMU-specific operations
├── LXC       - LXC-specific operations
├── Storage   - Storage management
├── Tasks     - Task monitoring
├── Auth      - Authentication operations
└── Version   - Version information
```

## Authentication

### Password Authentication

```go
client, err := pve.NewClient(
    "https://pve.example.com:8006",
    &pve.AuthOptions{
        Username: "root@pam",
        Password: "your-password",
        AuthType: pve.PasswordAuth,
    },
)
```

### API Token Authentication

```go
client, err := pve.NewClient(
    "https://pve.example.com:8006",
    &pve.AuthOptions{
        Username:    "root@pam",
        TokenID:     "mytoken",
        TokenSecret: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        AuthType:    pve.TokenAuth,
    },
)
```

## TLS Configuration

### Skip TLS Verification (Testing Only)

If you're using a Proxmox VE instance with self-signed certificates (common in development/testing), you can skip TLS verification:

```go
client, err := pve.NewClient(
    "https://pve.example.com:8006",
    &pve.AuthOptions{
        Username: "root@pam",
        Password: "your-password",
    },
    pve.WithInsecureTLS(), // Skip TLS certificate verification
)
```

**⚠️ WARNING**: Only use `WithInsecureTLS()` in testing or development environments. This makes connections vulnerable to man-in-the-middle attacks. In production, always use properly signed certificates or import the CA certificate to your system trust store.

## API Coverage

### Cluster Service (6 methods)

```go
client.Cluster.Get()                    // Get cluster information
client.Cluster.Status()                 // Get cluster status
client.Cluster.Resources()              // Get all cluster resources
client.Cluster.ResourcesByType(type)    // Get resources by type
client.Cluster.Tasks()                  // Get cluster tasks
client.Cluster.Nodes()                  // Get cluster nodes
```

### Nodes Service (23 methods)

**Basic Operations:**
```go
client.Nodes.List()                     // List all nodes
client.Nodes.Get(name)                  // Get node by name
client.Nodes.GetDetailed(name)          // Get detailed node info
client.Nodes.GetStatus(name)            // Get node status
client.Nodes.GetVersion(name)           // Get node version
client.Nodes.GetConfig(name)            // Get node configuration
```

**Lifecycle Management:**
```go
client.Nodes.Start(name)                // Start node
client.Nodes.Stop(name)                 // Stop node
client.Nodes.Shutdown(name)             // Shutdown node
client.Nodes.Reboot(name)               // Reboot node
```

**Monitoring & Resources:**
```go
client.Nodes.GetNetstat(name)           // Get network statistics
client.Nodes.GetSyslog(name, lines)     // Get system logs
client.Nodes.GetRRD(name, timeframe)    // Get RRD monitoring data
client.Nodes.GetNodeTasks(name, opts)   // Get node tasks
client.Nodes.GetStorage(name)           // Get node storage
client.Nodes.GetVMs(name)               // Get all VMs on node
client.Nodes.GetQEMUVMs(name)           // Get QEMU VMs on node
client.Nodes.GetLXCContainers(name)     // Get LXC containers on node
```

**Backup & Access:**
```go
client.Nodes.CreateVZDumpBackup(name, opts)      // Create backup
client.Nodes.ExtractVZDumpConfig(name, volume)   // Extract backup config
client.Nodes.CreateVNCShell(name)                // Create VNC shell
client.Nodes.GetSubscription(name)               // Get subscription info
```

### VMs Service (Generic, 14 methods)

```go
client.VMs.List(options)                // List all VMs (QEMU + LXC)
client.VMs.Get(vmid)                    // Get VM by ID
client.VMs.GetVMResource(vmid)          // Get VM resource info
client.VMs.GetStatus(vmid)              // Get VM status
client.VMs.Start(vmid)                  // Start VM
client.VMs.Stop(vmid)                   // Stop VM
client.VMs.Shutdown(vmid)               // Shutdown VM
client.VMs.Reboot(vmid)                 // Reboot VM
client.VMs.Suspend(vmid)                // Suspend VM
client.VMs.Resume(vmid)                 // Resume VM
client.VMs.Delete(vmid)                 // Delete VM
client.VMs.GetConfig(vmid)              // Get VM config
client.VMs.UpdateConfig(vmid, config)   // Update VM config
client.VMs.Clone(vmid, newID, name)     // Clone VM
```

### QEMU Service (32 methods)

**Basic Operations:**
```go
client.QEMU.List(node)                  // List QEMU VMs
client.QEMU.Get(node, vmid)             // Get VM info
client.QEMU.GetStatus(node, vmid)       // Get VM status
client.QEMU.GetConfig(node, vmid)       // Get VM config
client.QEMU.UpdateConfig(node, vmid, config)  // Update config
```

**Lifecycle Management:**
```go
client.QEMU.Start(node, vmid)           // Start VM
client.QEMU.Stop(node, vmid)            // Stop VM
client.QEMU.Shutdown(node, vmid)        // Shutdown VM
client.QEMU.Reboot(node, vmid)          // Reboot VM
client.QEMU.Reset(node, vmid)           // Hard reset VM (QEMU only)
client.QEMU.Suspend(node, vmid)         // Suspend VM
client.QEMU.Resume(node, vmid)          // Resume VM
client.QEMU.Delete(node, vmid)          // Delete VM
```

**Advanced Operations:**
```go
client.QEMU.Migrate(node, vmid, target, opts)     // Migrate VM
client.QEMU.Clone(node, vmid, newID, name, full)  // Clone VM
client.QEMU.ResizeDisk(node, vmid, disk, size)    // Resize disk
```

**Snapshot Management:**
```go
client.QEMU.ListSnapshots(node, vmid)                      // List snapshots
client.QEMU.CreateSnapshot(node, vmid, name, desc, state)  // Create snapshot (with VM state)
client.QEMU.DeleteSnapshot(node, vmid, snapName)           // Delete snapshot
client.QEMU.RollbackSnapshot(node, vmid, snapName)         // Rollback snapshot
```

**QEMU-Specific Features:**
```go
client.QEMU.SendMonitorCommand(node, vmid, cmd)   // QEMU monitor command
client.QEMU.GetVNCProxy(node, vmid, websocket)    // Get VNC proxy
```

**Guest Agent Operations:**
```go
client.QEMU.GetAgentInfo(node, vmid)                     // Get agent info
client.QEMU.GetAgentNetworkInterfaces(node, vmid)        // Get network interfaces
client.QEMU.GetAgentFilesystemInfo(node, vmid)           // Get filesystem info
client.QEMU.ExecuteAgentCommand(node, vmid, command)     // Execute command in guest
client.QEMU.GetAgentExecStatus(node, vmid, pid)          // Get command execution status
```

### LXC Service (27 methods)

**Basic Operations:**
```go
client.LXC.List(node)                   // List LXC containers
client.LXC.Get(node, vmid)              // Get container info
client.LXC.GetStatus(node, vmid)        // Get container status
client.LXC.GetConfig(node, vmid)        // Get container config
client.LXC.UpdateConfig(node, vmid, config)  // Update config
```

**Lifecycle Management:**
```go
client.LXC.Start(node, vmid)            // Start container
client.LXC.Stop(node, vmid)             // Stop container
client.LXC.Shutdown(node, vmid)         // Shutdown container
client.LXC.Reboot(node, vmid)           // Reboot container
client.LXC.Suspend(node, vmid)          // Suspend container
client.LXC.Resume(node, vmid)           // Resume container
client.LXC.Delete(node, vmid)           // Delete container
```

**Advanced Operations:**
```go
client.LXC.Migrate(node, vmid, target, opts)         // Migrate container
client.LXC.Clone(node, vmid, newID, hostname, full)  // Clone container
client.LXC.ResizeDisk(node, vmid, disk, size)        // Resize disk
```

**Snapshot Management:**
```go
client.LXC.ListSnapshots(node, vmid)              // List snapshots
client.LXC.CreateSnapshot(node, vmid, name, desc) // Create snapshot
client.LXC.DeleteSnapshot(node, vmid, snapName)   // Delete snapshot
client.LXC.RollbackSnapshot(node, vmid, snapName) // Rollback snapshot
```

**LXC-Specific Features:**
```go
client.LXC.GetInterfaces(node, vmid)     // Get container network interfaces
client.LXC.EnterContainer(node, vmid)    // Enter container shell
client.LXC.GetPending(node, vmid)        // Get pending config changes
client.LXC.GetVNCProxy(node, vmid, ws)   // Get VNC proxy
```

### Storage Service (10 methods)

```go
client.Storage.List(options)            // List all storage
client.Storage.Get(name)                // Get storage by name
client.Storage.GetContent(name)         // Get storage content
client.Storage.GetContentByType(name, type)  // Get content by type
client.Storage.ListContent(name, opts)  // List content with options
client.Storage.Upload(name, file, data) // Upload file
client.Storage.Download(name, volume)   // Download file
client.Storage.DeleteContent(name, vol) // Delete content
client.Storage.GetDir(name)             // Get directory listing
client.Storage.GetRRD(name, timeframe)  // Get RRD data
```

### Tasks Service (9 methods)

```go
client.Tasks.List(options)              // List all tasks
client.Tasks.GetTask(upid)              // Get task by UPID
client.Tasks.StopTask(upid)             // Stop task
client.Tasks.StopNodeTask(node, upid)   // Stop node task
client.Tasks.GetTaskLog(upid)           // Get task log
client.Tasks.GetTaskLogWithPaging(upid, start, limit)  // Get task log with paging
client.Tasks.GetNodeTaskLog(node, upid)           // Get node task log
client.Tasks.GetNodeTaskStatus(node, upid)        // Get node task status
client.Tasks.WaitForTask(upid, timeout)           // Wait for task completion
```

### Version Service (3 methods)

```go
client.Version.Get()                    // Get version information
client.Version.GetAPT()                 // Get APT version info
client.Version.GetPackages()            // Get available packages
```

### Auth Service (8 methods)

```go
client.Auth.Login(username, password)   // Login
client.Auth.Logout()                    // Logout
client.Auth.GetTicketInfo()             // Get ticket info
client.Auth.GetPermissions(path)        // Get permissions
client.Auth.GetUsers()                  // Get all users
client.Auth.GetUser(userid)             // Get specific user
client.Auth.CreateUser(userid, pass, email)  // Create user
client.Auth.DeleteUser(userid)          // Delete user
```

## Usage Examples

### QEMU Virtual Machines

```go
// List all QEMU VMs
vms, err := client.QEMU.List("pve-node1")
if err != nil {
    log.Fatal(err)
}

for _, vm := range vms {
    fmt.Printf("QEMU VM %d: %s (%s)\n", vm.ID, vm.Name, vm.Status)
}

// Get VM status
status, err := client.QEMU.GetStatus("pve-node1", 100)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("CPU: %.2f%%, Memory: %d MB / %d MB\n",
    status.CPU*100, status.Mem/(1024*1024), status.MaxMem/(1024*1024))

// Send QEMU monitor command (QEMU-specific)
result, err := client.QEMU.SendMonitorCommand("pve-node1", 100, "info version")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("QEMU Version: %s\n", result)

// Create snapshot with VM state
task, err := client.QEMU.CreateSnapshot("pve-node1", 100, "pre-update", "Before system update", true)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Snapshot task created: %s\n", task.UPID)

// Execute command via Guest Agent
exec, err := client.QEMU.ExecuteAgentCommand("pve-node1", 100, []string{"df", "-h"})
if err != nil {
    log.Fatal(err)
}

// Wait and get execution result
time.Sleep(1 * time.Second)
result, err := client.QEMU.GetAgentExecStatus("pve-node1", 100, exec.PID)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Command output:\n%s\n", result.OutData)

// Migrate VM to another node
migrateOpts := &pve.MigrateOptions{
    Online:           true,
    BWLimit:          100000,  // 100 MB/s
    MigrationNetwork: "10.0.0.0/24",
}
task, err = client.QEMU.Migrate("pve-node1", 100, "pve-node2", migrateOpts)
```

### LXC Containers

```go
// List all LXC containers
containers, err := client.LXC.List("pve-node1")
if err != nil {
    log.Fatal(err)
}

for _, ct := range containers {
    fmt.Printf("LXC Container %d: %s (%s)\n", ct.ID, ct.Name, ct.Status)
}

// Get container network interfaces (LXC-specific)
interfaces, err := client.LXC.GetInterfaces("pve-node1", 200)
if err != nil {
    log.Fatal(err)
}

for _, iface := range interfaces {
    fmt.Printf("Interface: %s\n", iface.Name)
    fmt.Printf("  MAC: %s\n", iface.HardwareAddress)
    for _, ip := range iface.IPAddresses {
        fmt.Printf("  IP: %s/%d (%s)\n", ip.IPAddress, ip.Prefix, ip.IPAddressType)
    }
}

// Enter container shell (LXC-specific)
termProxy, err := client.LXC.EnterContainer("pve-node1", 200)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Terminal available at port %v with ticket %s\n",
    termProxy["port"], termProxy["ticket"])

// Get pending configuration changes (LXC-specific)
pending, err := client.LXC.GetPending("pve-node1", 200)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Pending changes: %+v\n", pending)

// Clone container
task, err := client.LXC.Clone("pve-node1", 200, 201, "cloned-container", true)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Clone task: %s\n", task.UPID)
```

### Node Management

```go
// Get node status
status, err := client.Nodes.GetStatus("pve-node1")
if err != nil {
    log.Fatal(err)
}

if data, ok := status["data"].(map[string]interface{}); ok {
    fmt.Printf("Node: %s\n", "pve-node1")
    fmt.Printf("Status: %v\n", data["status"])
    fmt.Printf("Uptime: %v seconds\n", data["uptime"])

    if cpuinfo, ok := data["cpuinfo"].(map[string]interface{}); ok {
        fmt.Printf("CPU: %v cores (%v)\n", cpuinfo["cpus"], cpuinfo["model"])
    }

    if memory, ok := data["memory"].(map[string]interface{}); ok {
        total := memory["total"].(float64) / (1024*1024*1024)
        used := memory["used"].(float64) / (1024*1024*1024)
        fmt.Printf("Memory: %.2f GB / %.2f GB\n", used, total)
    }
}

// Get network statistics
netstat, err := client.Nodes.GetNetstat("pve-node1")
if err != nil {
    log.Fatal(err)
}

for _, conn := range netstat {
    fmt.Printf("%s %s -> %s (%s)\n",
        conn["proto"], conn["local_addr"], conn["remote_addr"], conn["state"])
}

// Create VZDump backup
backupAll := true
compress := "zstd"
remove := true

backupOpts := &pve.VZDumpOptions{
    All:              &backupAll,
    Mode:             "snapshot",
    Storage:          "backup-storage",
    Compress:         compress,
    Remove:           &remove,
    MaxFiles:         3,
    Mailto:           "admin@example.com",
    MailNotification: "failure",
    ZSTDThreads:      4,
}

task, err := client.Nodes.CreateVZDumpBackup("pve-node1", backupOpts)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Backup task created: %s\n", task.UPID)
```

### Task Monitoring

```go
// List all cluster tasks
tasks, err := client.Tasks.List(nil)
if err != nil {
    log.Fatal(err)
}

for _, task := range tasks {
    fmt.Printf("Task: %s - Type: %s, Status: %s, Node: %s\n",
        task.UPID, task.Type, task.Status, task.Node)
}

// Get specific task status
task, err := client.Tasks.GetTask("UPID:node1:00001234:00000000:5F123456:vzdump:100:root@pam:")
if err != nil {
    log.Fatal(err)
}

// Get task log
logs, err := client.Tasks.GetTaskLog(task.UPID)
if err != nil {
    log.Fatal(err)
}

for _, line := range logs {
    fmt.Println(line)
}

// Wait for task to complete
result, err := client.Tasks.WaitForTask(task.UPID, 300)  // 5 minutes timeout
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Task completed with status: %v\n", result["status"])
```

### Cluster Operations

```go
// Get cluster status
status, err := client.Cluster.Status()
if err != nil {
    log.Fatal(err)
}

for _, item := range status {
    if item.Type == "cluster" {
        fmt.Printf("Cluster: %s (Quorate: %d)\n", item.Name, item.Quorate)
    } else if item.Type == "node" {
        fmt.Printf("  Node: %s (Online: %d)\n", item.Name, item.Online)
    }
}

// Get all cluster resources
resources, err := client.Cluster.Resources()
if err != nil {
    log.Fatal(err)
}

for _, res := range resources {
    fmt.Printf("Resource: %s - Type: %s, Status: %s\n",
        res.Name, res.Type, res.Status)
}

// Get only QEMU VMs from cluster resources
qemuVMs, err := client.Cluster.ResourcesByType("qemu")
if err != nil {
    log.Fatal(err)
}

for _, vm := range qemuVMs {
    fmt.Printf("QEMU VM: %s on %s\n", vm.Name, vm.Node)
}
```

## Advanced Configuration

### Custom HTTP Client

```go
import (
    "net/http"
    "time"
)

httpClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}

client, err := pve.NewClient(
    "https://pve.example.com:8006",
    authOptions,
    pve.WithHTTPClient(httpClient),
)
```

### Custom Rate Limiting

```go
import "golang.org/x/time/rate"

// Allow 20 requests per second with burst of 5
limiter := rate.NewLimiter(rate.Limit(20), 5)

client, err := pve.NewClient(
    "https://pve.example.com:8006",
    authOptions,
    pve.WithRateLimiter(limiter),
)
```

### Custom User Agent

```go
client, err := pve.NewClient(
    "https://pve.example.com:8006",
    authOptions,
    pve.WithUserAgent("my-app/1.0.0"),
)
```

## Type Definitions

The library provides comprehensive type definitions for all API responses:

- `Node` - Proxmox node information
- `VM` - Virtual machine/container resource
- `VMStatus` - VM runtime status
- `VMConfig` - VM configuration
- `VMSnapshot` - VM snapshot information
- `Task` - Async task information
- `Storage` - Storage information
- `Cluster` - Cluster information
- `ClusterStatus` - Cluster status
- `ClusterResource` - Cluster resource
- `NetworkInterface` - Network interface info
- `NetworkIPAddress` - IP address info
- `FilesystemInfo` - Filesystem information
- `GuestAgent` - Guest agent information
- `GuestExec` - Guest execution info
- `GuestExecResult` - Guest execution result
- `MigrateOptions` - VM migration options
- `VZDumpOptions` - Backup options (30+ fields)

## Error Handling

The library provides structured error handling:

```go
vm, err := client.QEMU.Get("pve-node1", 100)
if err != nil {
    // Check for specific error types
    if strings.Contains(err.Error(), "401") {
        fmt.Println("Authentication error")
    } else if strings.Contains(err.Error(), "404") {
        fmt.Println("VM not found")
    } else {
        fmt.Printf("API error: %v\n", err)
    }
    return
}
```

## Testing

The library includes comprehensive tests. Run them with:

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Compatibility

- **Proxmox VE**: 8.x, 9.x
- **Go**: 1.24+

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Support

- 🐛 [Report Issues](https://github.com/ysicing/go-pve/issues)
- 📖 [API Documentation](https://pkg.go.dev/github.com/ysicing/go-pve)

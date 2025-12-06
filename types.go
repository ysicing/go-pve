package pve

import (
	"net/http"
)

// VM represents a Virtual Machine or Container
type VM struct {
	ID        int      `json:"vmid"`
	Name      string   `json:"name"`
	Node      string   `json:"node"`
	Type      string   `json:"type"`
	Status    string   `json:"status"`
	CPU       float64  `json:"cpu"`
	Mem       int64    `json:"mem"`
	MaxMem    int64    `json:"maxmem"`
	Disk      int64    `json:"disk"`
	MaxDisk   int64    `json:"maxdisk"`
	DiskRead  int64    `json:"diskread"`
	DiskWrite int64    `json:"diskwrite"`
	NetIn     int64    `json:"netin"`
	NetOut    int64    `json:"netout"`
	Uptime    int      `json:"uptime"`
	Template  int      `json:"template"`
	PID       int      `json:"pid"`
	Config    VMConfig `json:"-"`
	VMData    any      `json:"data"`
}

// VMListOptions specifies VM listing options
type VMListOptions struct {
	Content string `url:"content,omitempty"`
	VMType  string `url:"vmtype,omitempty"`
	Storage string `url:"storage,omitempty"`
	Node    string `url:"node,omitempty"`
	Enabled *bool  `url:"enabled,omitempty"`
	Full    *bool  `url:"full,omitempty"`
}

// StorageListOptions specifies storage listing options
type StorageListOptions struct {
	Storage   string `url:"storage,omitempty"`
	Content   string `url:"content,omitempty"`
	Enabled   *bool  `url:"enabled,omitempty"`
	StorageID string `url:"storageid,omitempty"`
}

// TaskListOptions specifies task listing options
type TaskListOptions struct {
	Since      string `url:"since,omitempty"`
	Source     string `url:"source,omitempty"`
	TaskFilter string `url:"task-filter,omitempty"`
}

// VMConfig holds VM configuration
type VMConfig struct {
	Name        string            `json:"name"`
	VMID        int               `json:"vmid"`
	Description string            `json:"description"`
	Cores       int               `json:"cores"`
	Memory      int               `json:"memory"`
	Storage     map[string]string `json:"storage"`
}

// VMs represents a list of VMs
type VMs []*VM

// Node represents a Proxmox node
type Node struct {
	Name       string  `json:"node"`
	Status     string  `json:"status"`
	CPU        float64 `json:"cpu"`
	MaxCPU     int     `json:"maxcpu"`
	Mem        int64   `json:"mem"`
	MaxMem     int64   `json:"maxmem"`
	Storage    int64   `json:"storage"`
	MaxStorage int64   `json:"maxstorage"`
	Uptime     int     `json:"uptime"`
	Level      string  `json:"level"`
	LocalTime  int     `json:"localtime"`
	Support    string  `json:"support"`
}

// Nodes represents a list of nodes
type Nodes []*Node

// Cluster represents cluster information
type Cluster struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Quorate     int    `json:"quorate"`
	Nodes       Nodes  `json:"nodes"`
	VersionInfo struct {
		Version string `json:"version"`
	} `json:"version"`
}

// ClusterResource represents a cluster resource
type ClusterResource struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Content string  `json:"content"`
	VMID    int     `json:"vmid"`
	Name    string  `json:"name"`
	Status  string  `json:"status"`
	MaxMem  int64   `json:"maxmem"`
	Mem     int64   `json:"mem"`
	MaxDisk int64   `json:"maxdisk"`
	Disk    int64   `json:"disk"`
	CPU     float64 `json:"cpu"`
	Uptime  int     `json:"uptime"`
	Node    string  `json:"node"`
	Plugin  string  `json:"plugin"`
	Storage string  `json:"storage"`
	Enable  int     `json:"enable"`
	Shared  int     `json:"shared"`
	Used    int64   `json:"used"`
	Avail   int64   `json:"avail"`
}

// ClusterResources represents a list of cluster resources
type ClusterResources []*ClusterResource

// ClusterStatus represents cluster status information
type ClusterStatus struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	ID      string `json:"id"`
	IP      string `json:"ip"`
	Online  int    `json:"online"`
	Level   string `json:"level"`
	Quorate int    `json:"quorate"`
	NodeID  int    `json:"nodeid"`
}

// Storage represents a storage entity
type Storage struct {
	Storage string `json:"storage"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Used    int64  `json:"used"`
	Avail   int64  `json:"avail"`
	Total   int64  `json:"total"`
	Enabled int    `json:"enabled"`
	Active  int    `json:"active"`
	Shared  int    `json:"shared"`
	Plugin  string `json:"plugin"`
	Nodes   string `json:"nodes"`
	ID      string `json:"id"`
}

// Storages represents a list of storage entities
type Storages []*Storage

// Task represents an async task
type Task struct {
	UPID      string `json:"upid"`
	ID        string `json:"id"`
	Node      string `json:"node"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	User      string `json:"user"`
	TokenID   string `json:"tokenid"`
	Saved     string `json:"saved"`
	StartTime int64  `json:"starttime"`
	EndTime   int64  `json:"endtime"`
	UPType    string `json:"type"`
	PID       int    `json:"pid"`
}

// Tasks represents a list of tasks
type Tasks []*Task

// Version represents version information
type Version struct {
	Version   string   `json:"version"`
	Reponame  string   `json:"repodate"`
	Release   string   `json:"release"`
	VersionID string   `json:"version_id"`
	Keyboard  string   `json:"keyboard"`
	Timezone  string   `json:"timezone"`
	Languages []string `json:"language"`
	PVE       struct {
		Version   string `json:"version"`
		Repoid    string `json:"repid"`
		Release   string `json:"release"`
		VersionID string `json:"version_id"`
	} `json:"pve"`
}

// Ticket represents an authentication ticket
type Ticket struct {
	CSRFPreventionToken string `json:"CSRFPreventionToken"`
	PVEAuthCookie       string `json:"PVEAuthCookie"`
	Ticket              string `json:"ticket"`
	Username            string `json:"username"`
}

// VMStatus represents VM status information
type VMStatus struct {
	VMID      int     `json:"vmid"`
	Name      string  `json:"name"`
	Status    string  `json:"status"`
	CPU       float64 `json:"cpu"`
	Mem       int64   `json:"mem"`
	MaxMem    int64   `json:"maxmem"`
	Disk      int64   `json:"disk"`
	MaxDisk   int64   `json:"maxdisk"`
	DiskRead  int64   `json:"diskread"`
	DiskWrite int64   `json:"diskwrite"`
	NetIn     int64   `json:"netin"`
	NetOut    int64   `json:"netout"`
	Uptime    int     `json:"uptime"`
	PID       int     `json:"pid"`
	HA        struct {
		Managed int `json:"managed"`
	} `json:"ha"`
	MaxProcs   int    `json:"maxprocs"`
	Config     string `json:"config"`
	CPUs       int    `json:"cpus"`
	QMPStatus  string `json:"qmpstatus"`
	Monitor    int    `json:"monitor"`
	Spice      int    `json:"spice"`
	SnapshotVM string `json:"snapshot"`
}

// VMSnapshot represents a VM snapshot
type VMSnapshot struct {
	SNAPSHOT string         `json:"sn"`
	VMID     int            `json:"vmid"`
	Name     string         `json:"name"`
	Date     int            `json:"date"`
	Parent   string         `json:"parent"`
	State    int            `json:"state"`
	Desc     string         `json:"description"`
	Config   map[string]any `json:"config"`
}

// VMSnapshots represents a list of VM snapshots
type VMSnapshots []*VMSnapshot

// GuestAgent represents guest agent information
type GuestAgent struct {
	Info struct {
		Version     string `json:"version"`
		GgTime      int    `json:"gg_time"`
		NetworkTime int    `json:"network_time"`
		HostTime    int    `json:"host_time"`
		TimeOffset  int    `json:"time-offset"`
	} `json:"info"`
	Network struct {
		Interfaces []GuestNetworkInterface `json:"interfaces"`
	} `json:"network"`
}

// GuestNetworkInterface represents a guest network interface
type GuestNetworkInterface struct {
	IPAddresses []GuestIPAddress `json:"ip-addresses"`
	MAC         string           `json:"mac"`
	Name        string           `json:"name"`
	Hardware    string           `json:"hardware"`
}

// GuestIPAddress represents an IP address
type GuestIPAddress struct {
	IP        string `json:"ip-address"`
	Netmask   string `json:"netmask"`
	IPVersion int    `json:"ip-version"`
	Scope     string `json:"scope"`
}

// GuestExec represents guest execution information
type GuestExec struct {
	PID       int    `json:"pid"`
	ExitCode  int    `json:"exit-code"`
	ExitTime  int    `json:"exit-time"`
	Directory string `json:"dir"`
	ENode     string `json:"enode"`
	Username  string `json:"username"`
	Stdout    int    `json:"stdout"`
	Stdin     int    `json:"stdin"`
	Stderr    int    `json:"stderr"`
}

// GuestExecResult represents guest execution result
type GuestExecResult struct {
	OutData  string `json:"out-data"`
	ErrData  string `json:"err-data"`
	Exited   int    `json:"exited"`
	ExitCode int    `json:"exit-code"`
}

// NodeInfo represents detailed node information
type NodeInfo struct {
	Node          string         `json:"node"`
	Cpuinfo       map[string]any `json:"cpuinfo"`
	Kversion      string         `json:"kversion"`
	MaxDisk       int64          `json:"maxdisk"`
	MaxMem        int64          `json:"maxmem"`
	SharedMaxDisk int64          `json:"shared_maxdisk"`
}

// RequestOptionFunc is a function type for request options
type RequestOptionFunc func(*http.Request) error

// WithHeader adds a header to the request
func WithHeader(key, value string) RequestOptionFunc {
	return func(req *http.Request) error {
		req.Header.Set(key, value)
		return nil
	}
}

// MigrateOptions specifies VM migration options
type MigrateOptions struct {
	Online           bool   // Online migration
	Force            bool   // Force migration
	MigrationNetwork string // Migration network
	BWLimit          int    // Bandwidth limit (KB/s)
	TargetStorage    string // Target storage
	Delete           bool   // Delete source data
}

// NetworkInterface represents a VM network interface
type NetworkInterface struct {
	Name             string              `json:"name"`
	HardwareAddress  string              `json:"hardware-address"`
	IPAddresses      []NetworkIPAddress  `json:"ip-addresses"`
	Statistics       *NetworkStatistics  `json:"statistics,omitempty"`
}

// NetworkIPAddress represents an IP address
type NetworkIPAddress struct {
	IPAddress     string `json:"ip-address"`
	IPAddressType string `json:"ip-address-type"`
	Prefix        int    `json:"prefix"`
}

// NetworkStatistics represents network statistics
type NetworkStatistics struct {
	RXBytes   int64 `json:"rx-bytes"`
	RXPackets int64 `json:"rx-packets"`
	RXErrors  int64 `json:"rx-errors"`
	RXDropped int64 `json:"rx-dropped"`
	TXBytes   int64 `json:"tx-bytes"`
	TXPackets int64 `json:"tx-packets"`
	TXErrors  int64 `json:"tx-errors"`
	TXDropped int64 `json:"tx-dropped"`
}

// FilesystemInfo represents filesystem information
type FilesystemInfo struct {
	Name        string `json:"name"`
	Mountpoint  string `json:"mountpoint"`
	Type        string `json:"type"`
	TotalBytes  int64  `json:"total-bytes"`
	UsedBytes   int64  `json:"used-bytes"`
	AvailBytes  int64  `json:"available-bytes,omitempty"`
}

// VZDumpOptions specifies vzdump backup options
type VZDumpOptions struct {
	VMID            string `url:"vmid,omitempty"`            // VM ID or list of IDs (comma-separated)
	All             *bool  `url:"all,omitempty"`             // Backup all VMs
	Mode            string `url:"mode,omitempty"`            // Backup mode: snapshot, suspend, stop
	Storage         string `url:"storage,omitempty"`         // Storage ID for backup
	Compress        string `url:"compress,omitempty"`        // Compression: 0, 1, gzip, lzo, zstd
	DumpDir         string `url:"dumpdir,omitempty"`         // Directory for backup files
	Remove          *bool  `url:"remove,omitempty"`          // Remove old backups
	MaxFiles        int    `url:"maxfiles,omitempty"`        // Maximum number of backup files
	Mailto          string `url:"mailto,omitempty"`          // Email address for notifications
	MailNotification string `url:"mailnotification,omitempty"` // Mail notification: always, failure
	Quiet           *bool  `url:"quiet,omitempty"`           // Suppress output
	Stop            *bool  `url:"stop,omitempty"`            // Stop mode
	Stopwait        int    `url:"stopwait,omitempty"`        // Max wait time for stop (minutes)
	Tmpdir          string `url:"tmpdir,omitempty"`          // Temporary directory
	Notes           string `url:"notes-template,omitempty"`  // Notes template
	Protected       *bool  `url:"protected,omitempty"`       // Protected backup
	PruneBackups    string `url:"prune-backups,omitempty"`   // Prune backups configuration
	Script          string `url:"script,omitempty"`          // Hook script
	Stdexcludes     *bool  `url:"stdexcludes,omitempty"`     // Exclude standard paths
	Stdout          *bool  `url:"stdout,omitempty"`          // Write to stdout
	Bwlimit         int    `url:"bwlimit,omitempty"`         // Bandwidth limit (KB/s)
	Ionice          int    `url:"ionice,omitempty"`          // IO priority (0-8)
	Lockwait        int    `url:"lockwait,omitempty"`        // Max wait time for lock (minutes)
	Performance     string `url:"performance,omitempty"`     // Performance settings
	Pigz            int    `url:"pigz,omitempty"`            // Pigz threads for compression
	Pool            string `url:"pool,omitempty"`            // Backup pool
	ZSTDThreads     int    `url:"zstd,omitempty"`            // ZSTD compression threads
}

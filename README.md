# go-pve

A Go client library for the Proxmox VE (PVE) API.

This library provides a simple and uniform way to interact with Proxmox VE's REST API from Go applications.

## Features

- Full support for Proxmox VE API endpoints
- Automatic authentication and token management
- Retry logic with exponential backoff
- Rate limiting support
- Comprehensive type definitions
- Clean, idiomatic Go API

## Installation

```bash
go get github.com/ysicing/go-pve
```

## Usage

### Basic Usage

```go
package main

import (
	"fmt"
	"log"

	pve "github.com/ysicing/go-pve"
)

func main() {
	client, err := pve.NewClient("https://pve.example.com:8006", &pve.AuthOptions{
		Username: "root@pam",
		Password: "your-password",
	})
	if err != nil {
		log.Fatal(err)
	}

	cluster, err := client.Cluster.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cluster: %s\n", cluster.Name)
	fmt.Printf("Nodes: %d\n", len(cluster.Nodes))
}
```

### Authentication

The library supports both password and API token authentication:

#### Password Authentication

```go
client, err := pve.NewClient("https://pve.example.com:8006", &pve.AuthOptions{
	Username: "root@pam",
	Password: "your-password",
})
```

#### API Token Authentication

```go
client, err := pve.NewClient("https://pve.example.com:8006", &pve.AuthOptions{
	Username: "root@pam",
	TokenID:  "your-token-id",
	TokenSecret: "your-token-secret",
})
```

### Working with VMs

```go
// List all VMs
vms, err := client.VMs.List()
if err != nil {
	log.Fatal(err)
}

for _, vm := range vms {
	fmt.Printf("VM %d: %s (%s)\n", vm.ID, vm.Name, vm.Status)
}

// Get VM details
vm, err := client.VMs.Get(100)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("VM: %s\n", vm.Name)

// Start VM
err = client.VMs.Start(100)
if err != nil {
	log.Fatal(err)
}
```

### Working with Nodes

```go
// List all nodes
nodes, err := client.Nodes.List()
if err != nil {
	log.Fatal(err)
}

for _, node := range nodes {
	fmt.Printf("Node: %s - Status: %s\n", node.Name, node.Status)
}
```

## Status

This library is in active development. The API may change until version 1.0.0 is released.

## License

MIT License - see LICENSE file for details.

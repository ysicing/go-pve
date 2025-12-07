// +build ignore

package main

import (
	"fmt"
	"log"

	pve "github.com/ysicing/go-pve"
)

func main() {
	// Create client with insecure TLS (skip certificate verification)
	// This is useful for development/testing with self-signed certificates
	client, err := pve.NewClient(
		"https://100.90.80.11:8006",
		&pve.AuthOptions{
			Username: "root@pam",
			Password: "your-password",
		},
		pve.WithInsecureTLS(), // Skip TLS certificate verification
	)
	if err != nil {
		log.Fatal(err)
	}

	// Test connection by getting cluster information
	version, err := client.Version.Get()
	if err != nil {
		log.Fatalf("Failed to get version: %v", err)
	}

	fmt.Printf("Connected to Proxmox VE %s\n", version.Version)
	fmt.Printf("Release: %s\n", version.Release)

	// Now you can safely use the client without certificate errors
	// Example: List nodes
	nodes, err := client.Nodes.List()
	if err != nil {
		log.Fatalf("Failed to list nodes: %v", err)
	}

	fmt.Printf("\nFound %d nodes:\n", len(nodes))
	for _, node := range nodes {
		fmt.Printf("  - %s (%s)\n", node.Name, node.Status)
	}
}

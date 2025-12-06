package main

import (
	"fmt"
	"log"

	pve "github.com/ysicing/go-pve"
)

func main() {
	// Create a new client with password authentication
	client, err := pve.NewClient("https://pve.example.com:8006", &pve.AuthOptions{
		Username: "root@pam",
		Password: "your-password",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Get cluster information
	cluster, err := client.Cluster.Get()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cluster Name: %s\n", cluster.Name)
	fmt.Printf("Cluster Version: %s\n", cluster.Version)
	fmt.Printf("Number of Nodes: %d\n", len(cluster.Nodes))

	// List all VMs
	vms, err := client.VMs.List(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nVMs on cluster:\n")
	for _, vm := range vms {
		fmt.Printf("  VM %d: %s (%s) - Node: %s, Status: %s\n",
			vm.ID, vm.Name, vm.Type, vm.Node, vm.Status)
	}

	// Get detailed info for a specific VM
	if len(vms) > 0 {
		vm, err := client.VMs.Get(vms[0].ID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\nVM %d Details:\n", vm.ID)
		fmt.Printf("  CPU: %.2f%%\n", vm.CPU*100)
		fmt.Printf("  Memory: %d MB / %d MB\n", vm.Mem/1024/1024, vm.MaxMem/1024/1024)
		fmt.Printf("  Disk: %d GB / %d GB\n", vm.Disk/1024/1024/1024, vm.MaxDisk/1024/1024/1024)
		fmt.Printf("  Uptime: %d seconds\n", vm.Uptime)
	}

	// List storage
	storages, err := client.Storage.List(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nStorage:\n")
	for _, storage := range storages {
		fmt.Printf("  %s: %d%% used (%d GB / %d GB)\n",
			storage.Storage,
			int(float64(storage.Used)/float64(storage.Total)*100),
			storage.Avail/1024/1024/1024,
			storage.Total/1024/1024/1024)
	}
}

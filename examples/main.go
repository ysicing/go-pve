package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	pve "github.com/ysicing/go-pve"
)

func main() {
	// Create a new client with password authentication
	client, err := pve.NewClient("https://100.90.80.11:8006", &pve.AuthOptions{
		Username:    "root@pam",
		TokenID:     "test",
		TokenSecret: "2c38303d-fccd-4d8b-a4a2-d74b9a5508dd",
	}, pve.WithInsecureTLS())
	if err != nil {
		log.Fatal(err)
	}

	nodes, err := client.Nodes.List()
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(nodes)
	for _, node := range nodes {
		fmt.Printf("Node: %s, Status: %s\n", node.Name, node.Status)
	}
}

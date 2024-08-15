package main

import (
	"fmt"
	"github.com/startcodextech/octopuslb/internal/logs"
	"github.com/startcodextech/octopuslb/pkg/dhcp"
)

func init() {
	logs.Init()
}

func main() {
	ifaces, err := dhcp.GetNetworkInterfaces()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	for _, iface := range ifaces {
		fmt.Printf("%+v\n", iface)
	}
}

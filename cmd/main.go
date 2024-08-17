package main

import (
	"github.com/phuslu/log"
	"github.com/startcodextech/octopuslb/internal/logs"
	"github.com/startcodextech/octopuslb/pkg/dhcp"
)

func init() {
	logs.Init()
}

func main() {
	interfaces, err := dhcp.GetNetworkInterfaces()
	if err != nil {
		log.Error().Err(err).Msg("failed to get network interfaces")
	}

	for _, iface := range interfaces {
		log.Printf("%+v", iface)
	}
}

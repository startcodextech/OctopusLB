package config

import "github.com/startcodextech/octopuslb/pkg/dhcp"

type DHCPConfig struct {
	State            *dhcp.State            `json:"server"`
	NetworkInterface *dhcp.NetworkInterface `json:"network_interface"`
}

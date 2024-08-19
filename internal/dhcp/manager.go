package dhcp

import (
	"github.com/startcodextech/octopuslb/internal/config"
	"github.com/startcodextech/octopuslb/pkg/dhcp"
)

type DHCPManager struct {
	server    *dhcp.DHCPServer
	config    config.DHCPConfig
	isStarted bool
}

func NewDHCPManager() (*DHCPManager, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	return &DHCPManager{
		config: cfg.DHCP,
	}, nil
}

func (m *DHCPManager) GetNetworksInterfaces() ([]dhcp.NetworkInterface, error) {
	return dhcp.GetNetworkInterfaces()
}

func (m *DHCPManager) ConfigureDHCP() error {
	server, err := dhcp.NewServer()
	if err != nil {
		return err
	}
	m.server = server
	return nil
}

func (m *DHCPManager) ChangeNetworkInterface(name string) error {
	return m.server.ChangeInterface(name)
}

func (m *DHCPManager) Start() error {
	err := m.server.Start()
	if err != nil {
		return err
	}
	m.isStarted = true
	return nil
}

func (m *DHCPManager) Stop() error {
	err := m.server.Stop()
	if err != nil {
		return err
	}
	m.isStarted = false
	return nil
}

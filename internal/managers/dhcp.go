package managers

import (
	"errors"
	"fmt"
	"github.com/startcodextech/octopuslb/internal/config"
	"github.com/startcodextech/octopuslb/pkg/dhcp"
	"net"
	"time"
)

var (
	ErrNotFoundNetworkInterface = errors.New("network interface not found")
)

type (
	DHCPManager struct {
		server       *dhcp.DHCPServer
		config       *config.DHCPConfig
		isStarted    bool
		isConfigured bool
	}

	DHCPConfig struct {
		InterfaceName string    `json:"interface_name"`
		IP            string    `json:"ip"`
		Mask          int       `json:"mask"`
		IPRange       [2]string `json:"ip_range"`
		DNS           []string  `json:"dns"`
		LeaseTime     int       `json:"lease_time"`
	}
)

func NewDHCPManager() (*DHCPManager, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	manager := &DHCPManager{
		config:       cfg.DHCP,
		isConfigured: true,
		isStarted:    false,
	}

	if cfg.DHCP == nil || cfg.DHCP.NetworkInterface == nil || cfg.DHCP.State == nil {
		manager.isConfigured = false
	}

	return manager, nil
}

func (m *DHCPManager) GetNetworksInterfaces() ([]dhcp.NetworkInterface, error) {
	return dhcp.GetNetworkInterfaces()
}

func (m *DHCPManager) ConfigureDHCP(setup DHCPConfig) error {
	var duration time.Duration
	var ip net.IP
	var mask net.IPMask
	var ipRange [2]net.IP
	var dns []net.IP
	if setup.LeaseTime == 0 {
		duration = dhcp.DefaultLeaseDuration
	}
	if setup.IP == "" {
		ip = dhcp.DefaultRouter
	}
	if setup.Mask == 0 {
		mask = dhcp.DefaultSubnetMask
	}
	if setup.IPRange[0] == "" {
		ipRange[0] = dhcp.DefaultRangeStartIp
	}
	if setup.IPRange[1] == "" {
		ipRange[1] = dhcp.DefaultRangeEndIp
	}
	if len(setup.DNS) == 0 {
		dns = dhcp.DefaultDNS
	} else {
		for _, d := range setup.DNS {
			dns = append(dns, net.ParseIP(d))
		}
	}

	state := &dhcp.State{
		Config: &dhcp.Config{
			LeaseDuration: duration,
			Router:        ip,
			SubnetMask:    mask,
			IPRange:       ipRange,
			DNS:           dns,
		},
		Leases: make(map[string]*dhcp.Lease),
	}

	netInterface, err := m.validateExistsNetInterface(setup.InterfaceName)
	if err != nil {
		return err
	}

	if m.server == nil {
		server, err := dhcp.NewServer(setup.InterfaceName, state)
		if err != nil {
			return fmt.Errorf("failed to create DHCP server: %w", err)
		}
		m.server = server
		return m.save()
	}
	m.server.Reload(state)
	err = m.server.ChangeInterface(setup.InterfaceName)
	if err != nil {
		return err
	}

	m.config = &config.DHCPConfig{
		State:            state,
		NetworkInterface: netInterface,
	}

	return m.save()
}

func (m *DHCPManager) validateExistsNetInterface(name string) (*dhcp.NetworkInterface, error) {
	interfaces, err := dhcp.GetNetworkInterfaces()
	if err != nil {
		return nil, fmt.Errorf("DHCP Manager failed get network adapters: %w", err)
	}

	for _, i := range interfaces {
		if i.Name == name || i.Alias == name {
			return &i, nil
		}
	}
	return nil, ErrNotFoundNetworkInterface
}

func (m *DHCPManager) save() error {
	configGlobal, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("Failed to load config on DHCP Manager: %w", err)
	}

	configGlobal.DHCP = m.config

	if err = config.SaveConfig(); err != nil {
		return fmt.Errorf("Failed to save config on DHCP Manager: %w", err)
	}
	return nil
}

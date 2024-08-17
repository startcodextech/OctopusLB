package dhcp

import (
	"errors"
	"net"
	"runtime"
)

type (
	NetworkInterface struct {
		Name       string   `json:"name"`
		Alias      string   `json:"alias,omitempty"`
		MacAddress string   `json:"mac_address"`
		IP         string   `ip:"ip"`
		Mask       string   `json:"mask"`
		Gateway    string   `json:"gateway"`
		DNS        []string `json:"dns"`
		Up         bool     `json:"up"`
	}
)

// GetNetworkInterfaces returns the names of the physical network interfaces
func GetNetworkInterfaces() ([]NetworkInterface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var macosMapping map[string]string
	if runtime.GOOS == "darwin" {
		macosMapping, err = getNetworkInterfaceMacOS()
		if err != nil {
			return nil, err
		}
	}
	var netInterfaces []NetworkInterface
	for _, iface := range interfaces {
		if isPhysicalInterface(iface) {
			var ip net.IP
			var mask net.IPMask
			var gatewayIp net.IP
			var dns []string
			adders, _ := iface.Addrs()
			for _, addr := range adders {
				if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
					ip = ipNet.IP
					mask = ipNet.Mask
				}
			}
			gateway, err := getIPv4Gateway(iface.Name)
			if err != nil {
				if !errors.Is(err, ErrNoGateway) {
					return nil, err
				}
			}
			if gateway != "" {
				gatewayIp = net.ParseIP(gateway)
			}
			dnsIPs, err := getDNS(iface.Name)
			if err != nil {
				if !errors.Is(err, ErrNotFountDNSForInterface) {
					return nil, err
				}
			}
			if len(dnsIPs) > 0 {
				dns = dnsIPs
			}

			if macosMapping != nil || len(macosMapping) > 0 {
				if service, ok := macosMapping[iface.Name]; ok {
					netInterfaces = append(netInterfaces, NetworkInterface{
						Name:       iface.Name,
						Alias:      service,
						MacAddress: iface.HardwareAddr.String(),
						Gateway:    gatewayIp.String(),
						IP:         ip.String(),
						Mask:       net.IP(mask).String(),
						DNS:        dns,
						Up:         iface.Flags&net.FlagUp != 0,
					})
				}
			} else {
				netInterface := NetworkInterface{
					Name:       iface.Name,
					MacAddress: iface.HardwareAddr.String(),
					Gateway:    gatewayIp.String(),
					IP:         ip.String(),
					Mask:       net.IP(mask).String(),
					DNS:        dns,
					Up:         iface.Flags&net.FlagUp != 0,
				}
				netInterfaces = append(netInterfaces, netInterface)
			}
		}
	}
	return netInterfaces, nil
}

func (i *NetworkInterface) getNameOrAlias() string {
	if runtime.GOOS == "darwin" {
		return i.Alias
	}
	return i.Name
}

func (i *NetworkInterface) ChangeState(up bool) error {
	state := "up"
	if !up {
		state = "down"
	}

	err := setInterfaceState(i.getNameOrAlias(), state)
	if err != nil {
		return err
	}

	i.Up = up

	return nil
}

func (i *NetworkInterface) ConfigureIPv4(gateway, ip string, mask int) error {
	err := setIPv4(i.getNameOrAlias(), gateway, ip, mask)
	if err != nil {
		return err
	}

	i.Gateway = gateway
	i.IP = ip
	i.Mask = net.IP(net.CIDRMask(mask, 32)).String()

	return nil
}

func (i *NetworkInterface) ConfigureDNS(dns []string) error {
	err := setDNS(i.getNameOrAlias(), dns)
	if err != nil {
		return err
	}

	i.DNS = dns

	return nil
}

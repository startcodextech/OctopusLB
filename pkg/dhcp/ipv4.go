package dhcp

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	INTERFACE_UP   = "up"
	INTERFACE_DOWN = "down"
)

var (
	toolsNetworkLinux   = []string{"ip"}
	toolsNetworkMacOS   = []string{"networksetup"}
	toolsNetworkWindows = []string{"netsh", "powershell"}

	excludedPrefixesLinux   = []string{"docker", "veth", "br-", "lo", "virbr", "wlxc", "tap", "tun", "vnet", "kube"}
	excludedPrefixesMacOS   = []string{"utun", "awdl", "llw", "vmnet", "gif", "stf", "p2p", "utap"}
	excludedPrefixesWindows = []string{"Loopback", "Virtual", "Teredo", "6to4", "isatap"}

	ErrNoGateway              = errors.New("no gateway found for interface")
	ErrUnsupportedToolNetwork = errors.New("unsupported network tool")
)

func getNetworkInterfaceMacOS() (map[string]string, error) {
	cmd := exec.Command("networksetup", "-listnetworkserviceorder")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	mapping := make(map[string]string)

	for _, line := range lines {
		if strings.Contains(line, "Hardware Port") {
			parts := strings.Split(line, ", Device: ")
			if len(parts) == 2 {
				device := strings.TrimSpace(parts[1])
				device = strings.TrimRight(device, ")")
				service := strings.Split(parts[0], ": ")[1]
				if device != "" {
					mapping[strings.TrimSpace(device)] = strings.TrimSpace(service)
				}
			}
		}
	}
	return mapping, nil
}

// GetNetworkInterfaces returns the names of the physical network interfaces
func GetNetworkInterfaces() ([]Interface, error) {
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
	var netInterfaces []Interface
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
					netInterfaces = append(netInterfaces, Interface{
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
				netInterface := Interface{
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

// getToolNetwork returns the path to the network tool (ifconfig, ip, netsh) based on the operating system
func getToolNetwork() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return getToolName(toolsNetworkLinux, ErrUnsupportedToolNetwork)
	case "darwin":
		return getToolName(toolsNetworkMacOS, ErrUnsupportedToolNetwork)
	case "windows":
		return getToolName(toolsNetworkWindows, ErrUnsupportedToolNetwork)
	default:
		return "", ErrUnsupportedToolNetwork
	}
}

// setInterfaceState sets the state of the network interface (up or down)
func setInterfaceState(name string, state string) error {
	networkTool, err := getToolNetwork()
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command(networkTool, "link", "set", name, state)
		return cmd.Run()
	case "darwin":
		on := "on"
		if state == "down" {
			on = "off"
		}
		cmd := exec.Command(networkTool, "-setnetworkserviceenabled", name, on)
		return cmd.Run()
	case "windows":
		switch networkTool {
		case "netsh":
			enable := "enable"
			if state == "down" {
				enable = "disable"
			}
			cmd := exec.Command(networkTool, "interface", "set", "interface", name, enable)
			return cmd.Run()
		case "powershell":
			tool := "Enable-NetAdapter"
			if state == "down" {
				tool = "Disable-NetAdapter"
			}
			cmd := exec.Command(tool, "-Name", name, "-Confirm:$false")
			return cmd.Run()
		}
		return ErrUnsupportedToolNetwork
	default:
		return ErrUnsupportedOS
	}
}

// isPhysicalInterface verify if the interface is a physical interface
func isPhysicalInterface(iface net.Interface) bool {
	if iface.Flags&net.FlagLoopback != 0 || len(iface.HardwareAddr) == 0 {
		return false
	}

	switch runtime.GOOS {
	case "darwin":
		excludedPrefixes := excludedPrefixesMacOS
		for _, prefix := range excludedPrefixes {
			if strings.HasPrefix(iface.Name, prefix) {
				return false
			}
		}
		return true //strings.HasPrefix(iface.Name, "en")
	case "linux":
		excludedPrefixes := excludedPrefixesLinux
		for _, prefix := range excludedPrefixes {
			if strings.HasPrefix(iface.Name, prefix) {
				return false
			}
		}
		return true
	case "windows":
		excludedPrefixes := excludedPrefixesWindows
		for _, prefix := range excludedPrefixes {
			if strings.Contains(iface.Name, prefix) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// configureInterface configures the network interface with the specified IP address and netmask
func configureIPv4(interfaceName, gateway, ipAddr string, mask int) error {
	tool, err := getToolNetwork()
	if err != nil {
		return err
	}

	ipMask, err := cidrToIP(mask)
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		return setIPv4Linux(tool, interfaceName, gateway, ipAddr, mask)
	case "darwin":
		return setIPv4MacOS(tool, interfaceName, gateway, ipAddr, ipMask)
	case "windows":
		return setIPv4Windows(tool, interfaceName, gateway, ipAddr, ipMask)
	default:
		return ErrUnsupportedOS
	}
}

func setIPv4Linux(tool, interfaceName, gateway, ipAddr string, mask int) error {
	cmd := exec.Command(tool, "addr", "add", ipAddr+"/"+strconv.Itoa(mask), "dev", interfaceName)
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command(tool, "route", "add", "default", "via", gateway, "dev", interfaceName)
	return cmd.Run()
}

func setIPv4MacOS(tool, interfaceName, gateway, ipAddr, ipMask string) error {
	cmd := exec.Command(tool, "-setmanual", interfaceName, ipAddr, ipMask)
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command(tool, "-setrouter", interfaceName, gateway)
	return cmd.Run()
}

func setIPv4Windows(tool, interfaceName, gateway, ipAddr, ipMask string) error {
	switch tool {
	case "netsh":
		// netsh interface ipv4 set address name="YOUR INTERFACE NAME" static IP_ADDRESS SUBNET_MASK GATEWAY
		// source=dhcp
		cmd := exec.Command(tool, "interface", "ipv4", "set", "address", `name="`+interfaceName+`"`, "static", ipAddr, ipMask, gateway)
		return cmd.Run()
	case "powershell":
		cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`New-NetIPAddress -InterfaceAlias "%s" -AddressFamily IPv4 -IPAddress %s -PrefixLength %s -DefaultGateway %s`, interfaceName, ipAddr, ipMask, gateway))
		return cmd.Run()
	default:
		return ErrUnsupportedToolNetwork
	}
}

// getIPv4Gateway returns the gateway IP address for the specified network interface
func getIPv4Gateway(interfaceName string) (string, error) {
	switch runtime.GOOS {
	case "linux":
		return getIPv4GatewayLinux(interfaceName)
	case "darwin":
		return getIPv4GatewayMac(interfaceName)
	case "windows":
		return getIPv4GatewayWindows(interfaceName)
	default:
		return "", ErrUnsupportedOS
	}
}

func getIPv4GatewayLinux(interfaceName string) (string, error) {
	cmd := fmt.Sprintf("ip route show dev %s | grep default | awk '{print $3}'", interfaceName)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	gateway := strings.TrimSpace(string(out))
	if gateway == "" {
		return "", ErrNoGateway
	}
	return gateway, nil
}

func getIPv4GatewayMac(interfaceName string) (string, error) {
	cmd := fmt.Sprintf("route -n get -ifscope %s default | grep gateway | awk '{print $2}'", interfaceName)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	gateway := strings.TrimSpace(string(out))
	if gateway == "" {
		return "", ErrNoGateway
	}
	return gateway, nil
}

func getIPv4GatewayWindows(interfaceName string) (string, error) {
	cmd := fmt.Sprintf(`Get-NetIPConfiguration -InterfaceAlias "%s" | Select-Object -ExpandProperty IPv4DefaultGateway | Select-Object -ExpandProperty NextHop`, interfaceName)
	out, err := exec.Command("powershell", "-Command", cmd).Output()
	if err != nil {
		return "", err
	}
	gateway := strings.TrimSpace(string(out))
	if gateway == "" {
		return "", ErrNoGateway
	}
	return gateway, nil
}

func cidrToIP(cidr int) (string, error) {
	if cidr < 0 || cidr > 32 {
		return "", fmt.Errorf("invalid CIDR value: %d", cidr)
	}
	mask := net.CIDRMask(cidr, 32)
	return net.IP(mask).String(), nil
}

func ipToCIDR(ipMask string) (int, error) {
	ip := net.ParseIP(ipMask)
	if ip == nil {
		return 0, fmt.Errorf("invalid IP address: %s", ipMask)
	}

	mask := net.IPMask(ip.To4())
	if mask == nil {
		return 0, fmt.Errorf("invalid IPv4 mask: %s", ipMask)
	}

	ones, _ := mask.Size()
	return ones, nil
}

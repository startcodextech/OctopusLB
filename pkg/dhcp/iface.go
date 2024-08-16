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
	toolsNetworkMacOS   = []string{"ifconfig"}
	toolsNetworkWindows = []string{"netsh", "powershell"}

	excludedPrefixesLinux   = []string{"docker", "veth", "br-", "lo", "virbr", "wlxc", "tap", "tun", "vnet", "kube"}
	excludedPrefixesMacOS   = []string{"utun", "awdl", "llw", "bridge", "vmnet", "gif", "stf", "p2p", "utap"}
	excludedPrefixesWindows = []string{"Loopback", "Virtual", "Teredo", "6to4", "isatap"}

	ErrNoGateway              = errors.New("no gateway found for interface")
	ErrUnsupportedToolNetwork = errors.New("unsupported network tool")
)

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

func getNetworkInterfaceMacOS() ([]Interface, error) {
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
				device = strings.TrimRight(device, ")") // Eliminar paréntesis de cierre si existen
				service := strings.Split(parts[0], ": ")[1]
				if device != "" { // Solo agregar si 'device' no está vacío
					mapping[device] = service
				}
			}
		}
	}

	for device, service := range mapping {
		var ip net.IP
		var mask net.IPMask
		var gatewayIp net.IP
		var dns []string

		iface := Interface{
			Name: device,
		}
	}
}

// GetNetworkInterfaces returns the names of the physical network interfaces
func GetNetworkInterfaces() ([]Interface, error) {

	os := runtime.GOOS

	switch os {
	case "linux", "windows":
		interfaces, err := net.Interfaces()
		if err != nil {
			return nil, err
		}
		var netInterfaces []Interface
		for _, iface := range interfaces {
			if isPhysicalInterface(iface) {
				var ip net.IP
				var mask net.IPMask
				var gatewayIp net.IP
				var dns []string
				addrs, _ := iface.Addrs()
				for _, addr := range addrs {
					if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
						ip = ipNet.IP
						mask = ipNet.Mask
					}
				}
				gateway, err := getGateway(iface.Name)
				if err != nil {

				}
				if gateway != "" {
					gatewayIp = net.ParseIP(gateway)
				}
				dnsIPs, err := getDNS(iface.Name)
				if err != nil {
				}
				if len(dnsIPs) > 0 {
					dns = dnsIPs
				}
				netIface := Interface{
					Name:       iface.Name,
					MacAddress: iface.HardwareAddr,
					Gateway:    gatewayIp,
					IP:         ip,
					Mask:       mask,
					DNS:        dns,
					Up:         iface.Flags&net.FlagUp != 0,
				}
				netInterfaces = append(netInterfaces, netIface)
			}
		}
		return netInterfaces, nil
	case "darwin":
		return getNetworkInterfaceMacOS()
	default:
		return nil, ErrUnsupportedOS
	}
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
		cmd := exec.Command(networkTool, name, state)
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

// configureInterface configures the network interface with the specified IP address and netmask
func configureInterface(interfaceName, gateway, ipAddr string, mask int) error {
	tool, err := getToolNetwork()
	if err != nil {
		return err
	}
	ipmask, err := cidrToIP(mask)
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command(tool, "addr", "add", ipAddr+"/"+strconv.Itoa(mask), "dev", interfaceName)
		err := cmd.Run()
		if err != nil {
			return err
		}
		cmd = exec.Command(tool, "route", "add", "default", "via", gateway, "dev", interfaceName)
		return cmd.Run()
	case "darwin":
		cmd := exec.Command(tool, interfaceName, ipAddr, "netmask", ipmask)
	case "windows":
		switch tool {
		case "netsh":
			// netsh interface ipv4 set address name="YOUR INTERFACE NAME" static IP_ADDRESS SUBNET_MASK GATEWAY
			// source=dhcp
			cmd = exec.Command(tool, "interface", "ipv4", "set", "address", `name="`+interfaceName+`"`, "static", ipAddr, mask, gateway)
		case "powershell":
			cmd = exec.Command("powershell", "-Command", fmt.Sprintf(`New-NetIPAddress -InterfaceAlias "%s" -AddressFamily IPv4 -IPAddress %s -PrefixLength %s -DefaultGateway %s`, interfaceName, ipAddr, mask, gateway))
		default:
			return ErrUnsupportedToolNetwork
		}

	}

	return cmd.Run()
}

// getGateway returns the gateway IP address for the specified network interface
func getGateway(interfaceName string) (string, error) {
	switch runtime.GOOS {
	case "linux":
		return getLinuxGateway(interfaceName)
	case "darwin":
		return getMacGateway(interfaceName)
	case "windows":
		return getWindowsGateway(interfaceName)
	default:
		return "", ErrUnsupportedOS
	}
}

func getLinuxGateway(interfaceName string) (string, error) {
	cmd := fmt.Sprintf("ip route show dev %s | grep default | awk '{print $3}'", interfaceName)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	gateway := strings.TrimSpace(string(out))
	if gateway == "" {
		return "", fmt.Errorf("no gateway found for interface %s", interfaceName)
	}
	return gateway, nil
}

func getMacGateway(interfaceName string) (string, error) {
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

func getWindowsGateway(interfaceName string) (string, error) {
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

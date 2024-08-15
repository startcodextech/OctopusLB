package dhcp

import (
	"errors"
	"fmt"
	"github.com/phuslu/log"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

var (
	excludedPrefixesLinux   = []string{"docker", "veth", "br-", "lo", "virbr", "wlxc", "tap", "tun", "vnet", "kube"}
	excludedPrefixesMacOS   = []string{"utun", "awdl", "llw", "bridge", "vmnet", "gif", "stf", "p2p", "utap"}
	excludedPrefixesWindows = []string{"Loopback", "Virtual", "Teredo", "6to4", "isatap"}

	ErrorNoGateway = errors.New("no gateway found for interface")
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

// GetNetworkInterfaces returns the names of the physical network interfaces
func GetNetworkInterfaces() ([]Interface, error) {
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
				log.Debug().
					Str("service", "dhcp").
					Str("interface", iface.Name).
					Err(err).
					Msg("No se pudo obtener la puerta de enlace")
			}
			if gateway != "" {
				gatewayIp = net.ParseIP(gateway)
			}
			dnsIPs, err := getDNS(iface.Name)
			if err != nil {
				log.Error().
					Str("service", "dhcp").
					Str("interface", iface.Name).
					Err(err).
					Msg("not able to get DNS servers")
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
			log.Debug().
				Str("service", "dhcp").
				Str("interface", iface.Name).
				Str("mac", netIface.MacAddress.String()).
				Str("ip", netIface.IP.String()).
				Str("mask", net.IP(netIface.Mask).String()).
				Str("gateway", netIface.Gateway.String()).
				Bool("up", netIface.Up).
				Msg("")
			netInterfaces = append(netInterfaces, netIface)
		}
	}
	return netInterfaces, nil
}

// getNetworkTool returns the path to the network tool (ifconfig, ip, netsh) based on the operating system
func getNetworkTool() (string, error) {
	switch runtime.GOOS {
	case "linux":
		ifconfigPath, err := exec.LookPath("ifconfig")
		if err == nil {
			return ifconfigPath, nil
		}

		ipPath, err := exec.LookPath("ip")
		if err == nil {
			return ipPath, nil
		}

		return "", errors.New("ifconfig and ip tools not found")
	case "darwin":
		ifconfigPath, err := exec.LookPath("ifconfig")
		if err == nil {
			return ifconfigPath, nil
		}
		return "", errors.New("ifconfig tool not found")
	case "windows":
		netshPath, err := exec.LookPath("netsh")
		if err == nil {
			return netshPath, nil
		}

		return "", errors.New("netsh tool not found")
	default:
		return "", errors.New("unsupported operating system")
	}
}

// setInterfaceState sets the state of the network interface (up or down)
func setInterfaceState(name string, state string) error {
	networkTool, err := getNetworkTool()
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "linux":
		if networkTool == "ifconfig" {
			cmd := exec.Command(networkTool, name, state)
			return cmd.Run()
		}
		cmd := exec.Command(networkTool, "link", "set", name, state)
		return cmd.Run()
	case "darwin":
		cmd := exec.Command(networkTool, name, state)
		return cmd.Run()
	case "windows":
		cmd := exec.Command(networkTool, "interface", "set", "interface", name, state)
		return cmd.Run()
	default:
		return errors.New("unsupported operating system")
	}
}

// configureInterface configures the network interface with the specified IP address and netmask
func configureInterface(ifaceName, ipAddr, mask string) error {
	tool, err := getNetworkTool()
	if err != nil {
		return err
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		if tool == "ifconfig" {
			cmd = exec.Command(tool, ifaceName, ipAddr, "netmask", mask)
		} else if tool == "ip" {
			cmd = exec.Command(tool, "addr", "add", ipAddr+"/"+mask, "dev", ifaceName)
		}
	case "darwin":
		cmd = exec.Command(tool, ifaceName, ipAddr, "netmask", mask)
	case "windows":
		cmd = exec.Command(tool, "interface", "ip", "set", "address", ifaceName, "static", ipAddr, mask)
	}

	return cmd.Run()
}

// getGateway returns the gateway IP address for the specified network interface
func getGateway(iface string) (string, error) {
	switch runtime.GOOS {
	case "linux":
		return getLinuxGateway(iface)
	case "darwin":
		return getMacGateway(iface)
	case "windows":
		return getWindowsGateway(iface)
	default:
		return "", errors.New("unsupported operating system")
	}
}

func getLinuxGateway(iface string) (string, error) {
	cmd := fmt.Sprintf("ip route show dev %s | grep default | awk '{print $3}'", iface)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	gateway := strings.TrimSpace(string(out))
	if gateway == "" {
		return "", fmt.Errorf("no gateway found for interface %s", iface)
	}
	return gateway, nil
}

func getMacGateway(iface string) (string, error) {
	cmd := fmt.Sprintf("route -n get -ifscope %s default | grep gateway | awk '{print $2}'", iface)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	gateway := strings.TrimSpace(string(out))
	if gateway == "" {
		return "", ErrorNoGateway
	}
	return gateway, nil
}

func getWindowsGateway(iface string) (string, error) {
	cmd := fmt.Sprintf(`Get-NetIPConfiguration -InterfaceAlias "%s" | Select-Object -ExpandProperty IPv4DefaultGateway | Select-Object -ExpandProperty NextHop`, iface)
	out, err := exec.Command("powershell", "-Command", cmd).Output()
	if err != nil {
		return "", err
	}
	gateway := strings.TrimSpace(string(out))
	if gateway == "" {
		return "", ErrorNoGateway
	}
	return gateway, nil
}

package dhcp

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var (
	toolsDnsLinux   = []string{"nmcli", "netconfig"}
	toolsDnsMacOS   = []string{"scutil", "networksetup"}
	toolsDnsWindows = []string{"netsh", "powershell"}

	ErrUnsupportedOS           = errors.New("unsupported operating system")
	ErrUnsupportedToolDNS      = errors.New("unsupported DNS tool")
	ErrNotFountDNSForInterface = errors.New("no DNS servers found for interface")
)

// setDNS configures the DNS servers for the specified network interface
func setDNS(interfaceName string, dnsServers []string) error {
	switch runtime.GOOS {
	case "linux":
		return setDNSLinux(interfaceName, dnsServers)
	case "darwin":
		return setDNSMacOS(interfaceName, dnsServers)
	case "windows":
		return setDNSWindows(interfaceName, dnsServers)
	default:
		return ErrUnsupportedOS
	}
}

func getDNS(interfaceName string) ([]string, error) {
	switch runtime.GOOS {
	case "linux":
		return getDNSLinux(interfaceName)
	case "darwin":
		return getDNSMacos(interfaceName)
	case "windows":
		return getDNSWindows(interfaceName)
	}
	return nil, ErrUnsupportedOS
}

// getToolDNS returns the DNS tool available on Linux
func getToolDNS() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return getToolName(toolsDnsLinux, ErrUnsupportedToolDNS)
	case "darwin":
		return getToolName(toolsDnsMacOS, ErrUnsupportedToolDNS)
	case "windows":
		return getToolName(toolsDnsWindows, ErrUnsupportedToolDNS)
	}
	return "", ErrUnsupportedOS
}

func getToolName(tools []string, err error) (string, error) {
	for _, tool := range tools {
		if _, err := exec.LookPath(tool); err == nil {
			return tool, nil
		}
	}
	return "", err
}

// setDNSLinux configures the DNS servers for the specified network interface on Linux
func setDNSLinux(interfaceName string, dnsServers []string) error {
	tool, err := getToolDNS()
	if err != nil {
		return err
	}

	dnsString := strings.Join(dnsServers, " ")

	switch tool {
	case "nmcli":
		cmd := exec.Command("nmcli", "con", "mod", interfaceName, "ipv4.dns", dnsString)
		err = cmd.Run()
		if err != nil {
			return err
		}
		return exec.Command("nmcli", "con", "up", interfaceName).Run()
	case "netconfig":
		cmd := exec.Command("sh", "-c", fmt.Sprintf("echo 'NETCONFIG_DNS_STATIC_SERVERS=\"%s\"' >> /etc/sysconfig/network/config", dnsString))
		if err := cmd.Run(); err != nil {
			return err
		}
		err := exec.Command("netconfig", "update", "-f").Run()
		if err != nil {
			return err
		}
		return nil
	}
	return ErrUnsupportedToolDNS
}

// setDNSMacOS configures the DNS servers for the specified network interface on macOS
func setDNSMacOS(interfaceName string, dnsServers []string) error {
	tool, err := getToolDNS()
	if err != nil {
		return err
	}

	switch tool {
	case "scutil":
		for _, dns := range dnsServers {
			cmd := exec.Command(tool, "--set", "DNS", dns)
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	case "networksetup":
		cmd := exec.Command("networksetup", "-setdnsservers", interfaceName, strings.Join(dnsServers, " "))
		return cmd.Run()
	}

	return ErrUnsupportedToolDNS
}

// setDNSWindows configures the DNS servers for the specified network interface on Windows
func setDNSWindows(interfaceName string, dnsServers []string) error {
	tool, err := getToolDNS()
	if err != nil {
		return err
	}

	switch tool {
	case "netsh":
		for i, dns := range dnsServers {
			index := fmt.Sprintf("index=%d", i+1)
			cmd := exec.Command("netsh", "interface", "ipv4", "set", "dns", `name="`+interfaceName+`"`, "static", dns, index)
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	case "powershell":
		cmd := exec.Command("powershell", "Set-DnsClientServerAddress", "-InterfaceAlias", interfaceName, "-ServerAddresses", strings.Join(dnsServers, ","))
		return cmd.Run()
	}

	return ErrUnsupportedToolDNS
}

func getDNSWindows(interfaceName string) ([]string, error) {
	tool, err := getToolDNS()
	if err != nil {
		return nil, err
	}

	var cmd *exec.Cmd
	switch tool {
	case "netsh":
		cmd = exec.Command("netsh", "interface", "ip", "show", "dns", interfaceName)
	case "powershell":
		psComd := fmt.Sprintf(`Get-DnsClientServerAddress -InterfaceAlias "%s" -AddressFamily IPv4 | Select-Object -ExpandProperty ServerAddresses`, interfaceName)
		cmd = exec.Command("powershell", "-Command", psComd)
	default:
		return nil, ErrUnsupportedToolDNS
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	output := string(out)

	var dnsServers []string
	ipRegex := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	matches := ipRegex.FindAllString(output, -1)

	for _, match := range matches {
		dnsServers = append(dnsServers, match)
	}

	if len(dnsServers) == 0 {
		return nil, ErrNotFountDNSForInterface
	}

	return dnsServers, nil
}

func getDNSMacos(interfaceName string) ([]string, error) {
	tool, err := getToolDNS()
	if err != nil {
		return nil, err
	}

	switch tool {
	case "scutil":
		cmd := exec.Command("scutil", "--dns")
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}

		var dnsServers []string
		var pendingDNS string

		scanner := bufio.NewScanner(strings.NewReader(string(out)))
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			if strings.HasPrefix(line, "resolver") {
				pendingDNS = ""
			}

			if strings.HasPrefix(line, "nameserver") {
				fields := strings.Fields(line)
				if len(fields) >= 3 {
					pendingDNS = fields[2]
				}
			}

			if strings.Contains(line, fmt.Sprintf("(%s)", interfaceName)) {
				if pendingDNS != "" && !contains(dnsServers, pendingDNS) {
					dnsServers = append(dnsServers, pendingDNS)
					pendingDNS = ""
				}
			}
		}

		if len(dnsServers) == 0 {
			return nil, ErrNotFountDNSForInterface
		}
		return dnsServers, nil
	}
	return nil, ErrUnsupportedToolDNS
}

func getDNSLinux(interfaceName string) ([]string, error) {
	tool, err := getToolDNS()
	if err != nil {
		return nil, err
	}

	switch tool {
	case "nmcli":
		cmd := exec.Command("nmcli", "dev", "show", interfaceName)
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}

		return parseDNSFromOutputLinux(string(out), "IP4.DNS")
	case "systemd-resolve":
		cmd := exec.Command("systemd-resolve", "--status", interfaceName)
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}

		return parseDNSFromOutputLinux(string(out), "DNS Servers")
	case "resolvectl":
		cmd := exec.Command("resolvectl", "status", interfaceName)
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		return parseDNSFromOutputLinux(string(out), "DNS Servers")
	case "dnsmasq":
		cmd := exec.Command("cat", "/var/run/dnsmasq/resolv.conf") // -> /etc/sysconfig/network/config
		out, err := cmd.Output()
		if err != nil {
			return nil, err
		}
		return parseDNSFromOutputLinux(string(out), "nameserver")
	default:
		return getDNSFromResolvConf()
	}
}

func parseDNSFromOutputLinux(output, prefix string) ([]string, error) {
	var dnsServers []string
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, prefix) {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				dnsServer := fields[len(fields)-1]
				if !contains(dnsServers, dnsServer) {
					dnsServers = append(dnsServers, dnsServer)
				}
			}
		}
	}

	if len(dnsServers) == 0 {
		return nil, ErrNotFountDNSForInterface
	}
	return dnsServers, nil
}

func getDNSFromResolvConf() ([]string, error) {
	file, err := os.Open("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var dnsServers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "nameserver") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				dnsServer := fields[1]
				if !contains(dnsServers, dnsServer) {
					dnsServers = append(dnsServers, dnsServer)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(dnsServers) == 0 {
		return nil, fmt.Errorf("no DNS servers found in /etc/resolv.conf")
	}
	return dnsServers, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

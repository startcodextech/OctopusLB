package nginx

import (
	"embed"
	"errors"
	"github.com/startcodextech/octopuslb/pkg/exec"
	"github.com/startcodextech/octopuslb/pkg/system"
	"os"
	"regexp"
	"strings"
)

//go:embed nginx.conf
var f embed.FS

var (
	ErrFailedVersion = errors.New("Failed to get Nginx version")
)

type Nginx struct {
	cmdInstall   exec.Instructions
	cmdUninstall exec.Instructions
	cmdEnable    exec.Instructions
	cmdDisable   exec.Instructions
	cmdReload    exec.Instructions
	cmdStart     exec.Instructions
	cmdStop      exec.Instructions
	cmdStatus    exec.Instructions
}

func New(os system.OSFamily) (*Nginx, error) {
	module := &Nginx{
		cmdInstall: exec.Instructions{
			{"apt", "update"},
			{"apt", "install", "-y", "curl", "gnupg2", "ca-certificates", "lsb-release", "openssl", "procps", "systemd", "ubuntu-keyring"},
			{"sh", "-c", "curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor | tee /usr/share/keyrings/nginx-archive-keyring.gpg >/dev/null"},
			{"sh", "-c", "gpg --dry-run --quiet --no-keyring --import --import-options import-show /usr/share/keyrings/nginx-archive-keyring.gpg"},
			{"sh", "-c", "echo \"deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] http://nginx.org/packages/ubuntu `lsb_release -cs` nginx\" | tee /etc/apt/sources.list.d/nginx.list"},
			{"sh", "-c", `echo -e "Package: *\nPin: origin nginx.org\nPin: release o=nginx\nPin-Priority: 900\n" | tee /etc/apt/preferences.d/99nginx`},
			{"apt", "update"},
			{"openssl", "dhparam", "-out", "/etc/ssl/certs/dhparam.pem", "2048"},
			{"apt", "install", "-y", "nginx"},
		},
		cmdUninstall: exec.Instructions{
			{"apt", "autoremove", "--purge", "-y", "nginx"},
			{"rm", "-f", "/etc/apt/sources.list.d/nginx.list"},
			{"rm", "-f", "/etc/apt/preferences.d/99nginx"},
			{"rm", "-f", "/usr/share/keyrings/nginx-archive-keyring.gpg"},
			{"apt", "update"},
		},
		cmdEnable: exec.Instructions{
			{"systemctl", "enable", "nginx"},
		},
		cmdStart: exec.Instructions{
			{"service", "nginx", "start"},
		},
		cmdDisable: exec.Instructions{
			{"systemctl", "disable", "nginx"},
		},
		cmdStop: exec.Instructions{
			{"nginx", "-s", "stop"},
		},
		cmdReload: exec.Instructions{
			{"nginx", "-s", "reload"},
		},
		cmdStatus: exec.Instructions{
			{"service", "nginx", "status"},
		},
	}

	switch os {
	case system.Debian:
		module.cmdInstall = exec.Instructions{
			{"apt", "update"},
			{"apt", "install", "-y", "curl", "gnupg2", "ca-certificates", "lsb-release", "systemd", "procps", "openssl", "debian-archive-keyring"},
			{"sh", "-c", "curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor | tee /usr/share/keyrings/nginx-archive-keyring.gpg > /dev/null"},
			{"sh", "-c", "gpg --dry-run --quiet --no-keyring --import --import-options import-show /usr/share/keyrings/nginx-archive-keyring.gpg"},
			{"sh", "-c", "echo \"deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] http://nginx.org/packages/debian `lsb_release -cs` nginx\" | tee /etc/apt/sources.list.d/nginx.list"},
			{"sh", "-c", `echo "Package: *\nPin: origin nginx.org\nPin: release o=nginx\nPin-Priority: 900" | tee /etc/apt/preferences.d/99nginx`},
			{"apt", "update"},
			{"openssl", "dhparam", "-out", "/etc/ssl/certs/dhparam.pem", "2048"},
			{"apt", "install", "-y", "nginx"},
		}
	case system.Rhel:
		module.cmdInstall = exec.Instructions{
			{"dnf", "makecache"},
			{"dnf", "install", "-y", "dnf-plugins-core", "systemd", "procps-ng", "openssl"},
			{"sh", "-c", `echo -e "[nginx-stable]\nname=nginx stable repo\nbaseurl=http://nginx.org/packages/centos/\$releasever/\$basearch/\ngpgcheck=1\nenabled=1\ngpgkey=https://nginx.org/keys/nginx_signing.key\nmodule_hotfixes=true\n\n[nginx-mainline]\nname=nginx mainline repo\nbaseurl=http://nginx.org/packages/mainline/centos/\$releasever/\$basearch/\ngpgcheck=1\nenabled=0\ngpgkey=https://nginx.org/keys/nginx_signing.key\nmodule_hotfixes=true" | tee /etc/yum.repos.d/nginx.repo`},
			{"dnf", "makecache"},
			{"openssl", "dhparam", "-out", "/etc/ssl/certs/dhparam.pem", "2048"},
			{"dnf", "install", "-y", "nginx"},
		}
		module.cmdUninstall = exec.Instructions{
			{"dnf", "remove", "-y", "nginx"},
			{"rm", "-f", "/etc/yum.repos.d/nginx.repo"},
			{"dnf", "autoremove", "-y"},
			{"dnf", "makecache"},
		}
	}

	return module, nil
}

// Install installs the Nginx service
func (n *Nginx) Install() error {
	for _, cmd := range n.cmdInstall {
		_, err := exec.RunCommand(cmd)
		if err != nil {
			return err
		}
	}

	err := n.copyFileConfig()
	if err != nil {
		return err
	}

	err = n.Enable()
	if err != nil {
		return err
	}

	err = n.Start()
	if err != nil {
		return err
	}

	return nil
}

// Uninstall uninstalls the Nginx service
func (n *Nginx) Uninstall() error {

	err := n.Stop()
	if err != nil {
		return err
	}

	err = n.Disable()
	if err != nil {
		return err
	}

	for _, cmd := range n.cmdUninstall {
		_, err = exec.RunCommand(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// Enable enables the Nginx service
func (n *Nginx) Enable() error {
	for _, cmd := range n.cmdEnable {
		_, err := exec.RunCommand(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// Disable disables the Nginx service
func (n *Nginx) Disable() error {
	for _, cmd := range n.cmdDisable {
		_, err := exec.RunCommand(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// Start starts the Nginx service
func (n *Nginx) Start() error {
	isRunning := n.IsRunning()

	if !isRunning {
		_, err := exec.RunCommand(n.cmdStart[0])
		if err != nil {
			return err
		}
	}

	return nil
}

// Stop stops the Nginx service
func (n *Nginx) Stop() error {
	isRunning := n.IsRunning()

	if isRunning {
		for _, cmd := range n.cmdStop {
			_, err := exec.RunCommand(cmd)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Reload reloads the Nginx configuration
func (n *Nginx) Reload() error {
	for _, cmd := range n.cmdReload {
		_, err := exec.RunCommand(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetVersion returns the Nginx version
func (n *Nginx) GetVersion() (string, error) {
	mgr, err := exec.DetectPackageManager()
	if err != nil {
		return "", err
	}

	_, err = exec.RunPkgMgr(exec.PkgMgrUpdate, "")
	if err != nil {
		return "", err
	}

	out, err := exec.RunPkgMgr(exec.PkgMgrInfo, "nginx")
	if err != nil {
		return "", err
	}

	var re *regexp.Regexp

	switch mgr {
	case exec.APT:
		re = regexp.MustCompile(`(?m)^Version:\s+(.+)$`)
	default:
		re = regexp.MustCompile(`(?m)^Version\s+:\s+(.+)$`)
	}

	outString := string(out)
	matches := re.FindStringSubmatch(outString)
	if len(matches) > 1 {
		return matches[1], nil
	}
	return "", ErrFailedVersion
}

// copyFileConfig copies the Nginx configuration file
func (n *Nginx) copyFileConfig() error {
	destinationPath := "/etc/nginx/nginx.conf"

	content, err := f.ReadFile("nginx.conf")
	if err != nil {
		return err
	}

	err = os.WriteFile(destinationPath, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

// IsRunning checks if Nginx is running
func (n *Nginx) IsRunning() bool {
	out, _ := exec.RunCommand(n.cmdStatus[0])
	if out != nil {
		if strings.TrimSpace(string(out)) == "nginx is running." {
			return true
		}
		return false
	}
	return false
}

package system

import (
	"errors"
)

var (
	ErrOSNotSupported = errors.New("OS not supported")
)

type (
	Package struct {
		cmdInstall   [][]string
		cmdUninstall [][]string
		cmdEnable    [][]string
		cmdDisable   [][]string
		cmdReload    [][]string
	}

	// nginx -> https://nginx.org/en/linux_packages.html
	Packages struct {
		Nginx Package
	}
)

func GetPackages(os string) (Packages, error) {
	pkgs := Packages{}

	switch os {
	case "ubuntu":
		pkgs.Nginx = Package{
			cmdInstall: [][]string{
				{"apt", "update"},
				{"apt", "install", "-y", "curl", "gnupg2", "ca-certificates", "lsb-release", "ubuntu-keyring"},
				{"sh", "-c", "curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor | tee /usr/share/keyrings/nginx-archive-keyring.gpg >/dev/null"},
				{"sh", "-c", "gpg --dry-run --quiet --no-keyring --import --import-options import-show /usr/share/keyrings/nginx-archive-keyring.gpg"},
				{"sh", "-c", "echo \"deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] http://nginx.org/packages/ubuntu `lsb_release -cs` nginx\" | tee /etc/apt/sources.list.d/nginx.list"},
				{"sh", "-c", `echo -e "Package: *\nPin: origin nginx.org\nPin: release o=nginx\nPin-Priority: 900\n" | tee /etc/apt/preferences.d/99nginx`},
				{"apt", "update"},
				{"apt", "install", "-y", "nginx"},
			},
			cmdUninstall: [][]string{
				{"apt", "autoremove", "--purge", "-y", "nginx"},
				{"rm", "-f", "/etc/apt/sources.list.d/nginx.list"},
				{"rm", "-f", "/etc/apt/preferences.d/99nginx"},
				{"rm", "-f", "/usr/share/keyrings/nginx-archive-keyring.gpg"},
				{"apt", "update"},
			},
		}
	case "debian":
		pkgs.Nginx = Package{
			cmdInstall: [][]string{
				{"apt", "update"},
				{"apt", "install", "-y", "curl", "gnupg2", "ca-certificates", "lsb-release", "debian-archive-keyring"},
				{"sh", "-c", "curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor | tee /usr/share/keyrings/nginx-archive-keyring.gpg > /dev/null"},
				{"sh", "-c", "gpg --dry-run --quiet --no-keyring --import --import-options import-show /usr/share/keyrings/nginx-archive-keyring.gpg"},
				{"sh", "-c", "echo \"deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] http://nginx.org/packages/debian `lsb_release -cs` nginx\" | tee /etc/apt/sources.list.d/nginx.list"},
				{"sh", "-c", `echo -e "Package: *\nPin: origin nginx.org\nPin: release o=nginx\nPin-Priority: 900" | tee /etc/apt/preferences.d/99nginx`},
				{"apt", "update"},
				{"apt", "install", "-y", "nginx"},
			},
			cmdUninstall: [][]string{
				{"apt", "autoremove", "--purge", "-y", "nginx"},
				{"rm", "-f", "/etc/apt/sources.list.d/nginx.list"},
				{"rm", "-f", "/etc/apt/preferences.d/99nginx"},
				{"rm", "-f", "/usr/share/keyrings/nginx-archive-keyring.gpg"},
				{"apt", "update"},
			},
		}
	case "rhel":
		pkgs.Nginx = Package{
			cmdInstall: [][]string{
				{"yum", "makecache"},
				{"yum", "install", "-y", "yum-utils"},
				{"sh", "-c", `echo -e "[nginx-stable]\nname=nginx stable repo\nbaseurl=http://nginx.org/packages/centos/\$releasever/\$basearch/\ngpgcheck=1\nenabled=1\ngpgkey=https://nginx.org/keys/nginx_signing.key\nmodule_hotfixes=true\n\n[nginx-mainline]\nname=nginx mainline repo\nbaseurl=http://nginx.org/packages/mainline/centos/\$releasever/\$basearch/\ngpgcheck=1\nenabled=0\ngpgkey=https://nginx.org/keys/nginx_signing.key\nmodule_hotfixes=true" | tee /etc/yum.repos.d/nginx.repo`},
				{"yum", "makecache"},
				{"yum", "install", "-y", "nginx"},
			},
			cmdUninstall: [][]string{
				{"yum", "remove", "-y", "nginx"},
				{"rm", "-f", "/etc/yum.repos.d/nginx.repo"},
				{"yum", "autoremove", "-y"},
				{"yum", "makecache"},
			},
		}
	default:
		return pkgs, ErrOSNotSupported
	}

	return pkgs, nil
}

package system

import (
	"errors"
)

var (
	ErrorPackageManagerNotFound = errors.New("Package not found")
)

type OS struct {
	Name            string   `json:"name"`
	Platform        string   `json:"platform"`
	PlatformFamily  string   `json:"platform_family"`
	PlatformVersion string   `json:"platform_version"`
	PackManager     string   `json:"-"`
	Packages        Packages `json:"-"`
}

func (os *OS) InstallPackage(name string) error {
	var err error
	switch name {
	case "nginx":
		for _, cmdLine := range os.Packages.Nginx.cmdInstall {
			_, err = runCommands(cmdLine)
			if err != nil {
				break
			}
		}
	default:
		err = ErrorPackageManagerNotFound
	}
	return err
}

func (os *OS) UninstallPackage(name string) error {
	var err error
	switch name {
	case "nginx":
		for _, cmdLine := range os.Packages.Nginx.cmdUninstall {
			_, err = runCommands(cmdLine)
			if err != nil {
				break
			}
		}
	default:
		err = ErrorPackageManagerNotFound
	}
	return err
}

func (os *OS) UpdateSystem() error {
	return runPackageManagerCommand("update", "")
}

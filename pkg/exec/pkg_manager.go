package exec

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var (
	managers = []PkgMgr{APT, DNF, ZYPPER, APK}

	commands = map[PkgMgrAction]map[PkgMgr]string{
		"install": {
			"apt":    "install -y",
			"dnf":    "install -y",
			"zypper": "install -y",
			"apk":    "add --no-cache",
		},
		"uninstall": {
			"apt":    "remove -y",
			"dnf":    "remove -y",
			"zypper": "remove -y",
			"apk":    "del",
		},
		"update": {
			"apt":    "update",
			"dnf":    "makecache",
			"zypper": "refresh",
			"apk":    "update",
		},
		"info": {
			"apt":    "show",
			"dnf":    "info",
			"zypper": "info",
			"apk":    "info",
		},
	}

	ErrPackageManagerNotFound      = errors.New("package manager not found")
	ErrInvalidActionForInstruction = errors.New("invalid action for instruction")
	ErrInvalidInstructionForAction = fmt.Errorf("invalid instruction %s for action %s")
)

type (
	PkgMgrAction string
	PkgMgr       string
)

const (
	APT    PkgMgr = "apt"
	DNF    PkgMgr = "dnf"
	ZYPPER PkgMgr = "zypper"
	APK    PkgMgr = "apk"

	PkgMgrInstall    PkgMgrAction = "install"
	PkgMgrUninstall  PkgMgrAction = "uninstall"
	PkgMgrUpdate     PkgMgrAction = "update"
	PkgMgrInfo       PkgMgrAction = "info"
	PkgMgrAutoremove PkgMgrAction = "autoremove"
)

func RunPkgMgr(action PkgMgrAction, packages string) ([]byte, error) {
	pkgs := strings.TrimSpace(packages)
	pm, err := DetectPackageManager()
	if err != nil {
		return nil, err
	}

	instruction, err := getInstruction(pm, action)
	if err != nil {
		return nil, err
	}

	args := append(strings.Split(instruction, " "), strings.Split(pkgs, " ")...)
	if pkgs == "" {
		args = strings.Split(instruction, " ")
	}

	cmd := exec.Command(string(pm), args...)
	out, err := cmd.Output()
	if err != nil {
		log.Println(pm, action, pkgs, "-> Failed")
		return nil, err
	}
	if pkgs == "" {
		log.Println(pm, action, "-> OK")
	} else {
		log.Println(pm, action, pkgs, "-> OK")
	}

	return out, nil
}

func DetectPackageManager() (PkgMgr, error) {
	for _, pm := range managers {
		if commandExists(string(pm)) {
			return pm, nil
		}
	}
	return "", ErrPackageManagerNotFound
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func getInstruction(pkgManager PkgMgr, action PkgMgrAction) (string, error) {
	if _, ok := commands[action]; !ok {
		return "", ErrInvalidActionForInstruction
	}

	if instruction, ok := commands[action][pkgManager]; ok {
		return instruction, nil
	}

	return "", ErrInvalidInstructionForAction
}

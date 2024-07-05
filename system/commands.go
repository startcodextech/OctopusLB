package system

import (
	"errors"
	"fmt"
	"github.com/phuslu/log"
	"os/exec"
	"strings"
)

var (
	managers = []string{"apt", "yum", "zypper", "apk"}

	commands = map[string]map[string]string{
		"install": {
			"apt":    "install -y",
			"yum":    "install -y",
			"zypper": "install -y",
			"apk":    "add --no-cache",
		},
		"uninstall": {
			"apt":    "remove -y",
			"yum":    "remove -y",
			"zypper": "remove -y",
			"apk":    "del",
		},
		"update": {
			"apt":    "update",
			"yum":    "makecache",
			"zypper": "refresh",
			"apk":    "update",
		},
	}

	ErrPackageManagerNotFound      = errors.New("package manager not found")
	ErrInvalidActionForInstruction = errors.New("invalid action for instruction")
	ErrInvalidInstructionForAction = fmt.Errorf("invalid instruction %s for action %s")
)

func runPackageManagerCommand(action, packages string) error {
	pm, err := detectPackageManager()
	if err != nil {
		log.Error().Err(err).Msg("Failed to detect package manager")
		return err
	}

	instruction, err := getInstruction(pm, action)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get instruction")
		return err
	}

	args := append(strings.Split(instruction, " "), strings.Split(packages, " ")...)
	if packages == "" {
		args = strings.Split(instruction, " ")
	}

	log.Info().
		Str("package_manager", pm).
		Str("instruction", instruction).
		Str("packages", packages).
		Msg("Running package manager command")

	cmd := exec.Command(pm, args...)
	_, err = cmd.Output()
	if err != nil {
		log.Error().Err(err).Msg("Failed to run package manager command")
		return err
	}
	log.Info().Str("packages", packages).Msgf("Packages %s used", action)
	return nil
}

func detectPackageManager() (string, error) {
	for _, pm := range managers {
		if commandExists(pm) {
			return pm, nil
		}
	}
	return "", ErrPackageManagerNotFound
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func getInstruction(pm, action string) (string, error) {
	if _, ok := commands[action]; !ok {
		log.Error().Err(ErrInvalidActionForInstruction).Msg("Invalid action")
		return "", ErrInvalidActionForInstruction
	}

	if instruction, ok := commands[action][pm]; ok {
		return instruction, nil
	}

	return "", ErrInvalidInstructionForAction
}

func runCommands(cmdLine []string) ([]byte, error) {
	if len(cmdLine) == 0 {
		log.Error().Str("command", "empty").Msg("❌")
		return nil, errors.New("command line is empty")
	}
	cmd := exec.Command(cmdLine[0])
	if len(cmdLine) > 1 {
		cmd = exec.Command(cmdLine[0], cmdLine[1:]...)
	}

	fullCommand := strings.Join(cmdLine, " ")

	out, err := cmd.Output()
	if err != nil {
		log.Error().Str("command", fullCommand).Err(err).Msg("❌")
		if len(out) > 0 {
			log.Printf("\n%s", out)
		}
		return nil, err
	}
	log.Info().Str("command", fullCommand).Msg("✅")
	if len(out) > 0 {
		log.Printf("\n%s", out)
	}
	return out, nil
}

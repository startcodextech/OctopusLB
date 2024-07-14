package exec

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

type (
	Instruction  []string
	Instructions []Instruction
)

func RunCommand(cmdLine []string) ([]byte, error) {
	var stdout, stderr *bytes.Buffer

	if len(cmdLine) == 0 {
		log.Println("Command line is empty")
		return nil, errors.New("command line is empty")
	}
	cmd := exec.Command(cmdLine[0])
	if len(cmdLine) > 1 {
		cmd = exec.Command(cmdLine[0], cmdLine[1:]...)
	}

	fullCommand := strings.Join(cmdLine, " ")

	if cmd.Stderr == nil {
		stderr = bytes.NewBuffer(make([]byte, 0, 1024))
		cmd.Stderr = stderr
	}
	if cmd.Stdout == nil {
		stdout = bytes.NewBuffer(make([]byte, 0, 1024))
		cmd.Stdout = stdout
	}

	err := cmd.Run()

	out := append(stdout.Bytes(), stderr.Bytes()...)

	if stderr != nil && err != nil {
		log.Println(fullCommand, "-> Failed")
		return out, err
	}
	log.Println(fullCommand, "-> OK")
	return out, nil
}

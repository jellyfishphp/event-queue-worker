package process

import (
	"os/exec"
	"strings"
)

type ProcessInterface interface {
	Start() ([]byte, error)
	IsRunning() bool
}

type Process struct {
	command   string
	isRunning bool
}

func (process *Process) Start() ([]byte, error) {
	process.isRunning = true

	commandParts := strings.Split(process.command, " ")

	command := exec.Command(commandParts[0], commandParts[1:]...)
	output, error := command.Output()

	process.isRunning = false

	return output, error
}

func (process *Process) IsRunning() bool {
	return process.isRunning
}

func NewProcess(command string) ProcessInterface {
	return &Process{
		command:   command,
		isRunning: false,
	}
}

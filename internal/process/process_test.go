package process

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestIsRunning(t *testing.T) {
	currentProcess := NewProcess("ls")

	if currentProcess.IsRunning() == true {
		t.Errorf("Process must not run.")
	}
}

func TestStart(t *testing.T) {
	expectedOutput := []byte("foo")
	command := fmt.Sprintf("echo '%s'", string(expectedOutput))

	currentProcess := NewProcess(command)

	output, err := currentProcess.Start()

	if err != nil {
		t.Errorf(err.Error())
	}

	if bytes.Equal(output, expectedOutput) {
		t.Errorf("Output must be '%s'", string(expectedOutput))
	}
}

func TestIsRunningAfterAsyncStart(t *testing.T) {
	currentProcess := NewProcess("sleep 2s")

	go func(process ProcessInterface) {
		process.Start()
	}(currentProcess)

	time.Sleep(time.Second)

	if currentProcess.IsRunning() == false {
		t.Errorf("Process must be running.")
	}
}

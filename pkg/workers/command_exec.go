package workers

import (
	"os"
	"os/exec"
)

// RunCommand executes a command with the given name and arguments.
func RunCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

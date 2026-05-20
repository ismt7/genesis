package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func runAttachedCommand(name string, args []string, actionName string) error {
	command := exec.Command(name, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to run %s: %w", actionName, err)
	}

	return nil
}

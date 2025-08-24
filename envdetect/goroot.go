package envdetect

import (
	"fmt"     // Format & Output Strings
	"os/exec" // Execute CLI commands
)

func FindGOROOT() (string, error) {
	goroot, gorootErr := exec.Command("go", "env", "GOROOT").Output() // Finds the directory of the Go installation (GOROOT)

	if gorootErr != nil {
		return "", fmt.Errorf("there was an error in finding your GOROOT: %w. \n Have you installed Go, and is it in your PATH?", gorootErr)
	}
	return string(goroot), nil
}

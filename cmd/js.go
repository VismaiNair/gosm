package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"       // Execute CLI commands
	"path/filepath" // Work with file paths
	"strings"       // Deal with strings

	"github.com/spf13/cobra"  // Importing the Cobra library for CLI apps
	"golang.org/x/mod/semver" // The semver library by Golang for comparing different versions
)

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open source file %s: %w", src, err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("could not create destination file %s: %w", dst, err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("could not copy file: %w", err)
	}
	return nil
}

// jsCmd represents the js command
var jsCmd = &cobra.Command{
	Use:   "js",
	Short: "Copies wasm_exec.js from Go installation so that WASM can be executed.",
	Long:  `The command gosm js copies wasm_exec.js from your go installation. This is required for running WASM modules in the browser if they were compiled from Go or Node.js on your machine.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		goroot, gorootErr := exec.Command("go", "env", "GOROOT").Output() // Finds the directory of the Go installation (GOROOT)

		if gorootErr != nil {
			return fmt.Errorf("there was an error in finding your GOROOT: %w. \n Have you installed Go, and is it in your PATH?", gorootErr)
		}

		fmt.Println("Finding Go Version...")
		version, versionErr := exec.Command("go", "version").Output() // Finds the version of Go
		fmt.Println("Go version found: " + string(version))

		if versionErr != nil {
			return fmt.Errorf("there was an error in finding your Go version: %w. \n Have you installed Go, and is it in your PATH?", versionErr)
		}
		// Split the string, get the third element, trim the "go" prefix, and add "v".
		cleanedVersion := "v" + strings.TrimPrefix(strings.Split(string(version), " ")[2], "go")
		fmt.Println("The cleaned go version is: " + cleanedVersion)

		var src string
		if semver.Compare(cleanedVersion, "v1.21") >= 0 {
			src = filepath.Join(strings.TrimSpace(string(goroot)), "lib", "wasm", "wasm_exec.js")
			fmt.Println("Go is greater than v1.21.")
		} else if semver.Compare(cleanedVersion, "v1.21") == -1 {
			src = filepath.Join(strings.TrimSpace(string(goroot)), "misc", "wasm", "wasm_exec.js")
		}

		if err := copyFile(src, "wasm_exec.js"); err != nil {
			return fmt.Errorf("there was an error with copying wasm_exec.js: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(jsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

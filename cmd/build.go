package cmd

import (
	"fmt"
	"os"      // For working with the operating syste
	"os/exec" // For executing cli commands

	"github.com/spf13/cobra" // Importing the Cobra library for CLI apps
)

var buildCmd = &cobra.Command{
	Use:   "build [package]",
	Short: "Build Go code to WebAssembly.",
	Long:  `Build Go code to WebAssembly (WASM) format. /nYou can specify a package or go file to build. If neither are specified, it defaults to all go files in the current directory. /nUnder the hood, it runs the command 'go build GOOS=js GOARCH=wasm'. This command is useful for preparing Go applications to run in a web environment using WebAssembly.`,
	Run: func(cmd *cobra.Command, args []string) {
		target := "." // Default target is current directory
		if len(args) > 0 {
			target = args[0] // If a package is specified, use it as target
		}

		buildCmd := exec.Command("go", "build", target) // Run the go build command
		fmt.Fprintf(os.Stdout, "Creating WASM Build command... \n")
		buildCmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm") // Set the environment variables for WebAssembly

		buildCmd.Stdout = os.Stdout // Redirect standard output to the console
		buildCmd.Stderr = os.Stderr // Redirect standard error to the console
		fmt.Fprintf(os.Stdout, "Building WASM from provided file or package... \n")
		buildCmd.Run() // Execute the build command
	},
}

func init() {
	rootCmd.AddCommand(buildCmd) // Add the build command to the root command
}

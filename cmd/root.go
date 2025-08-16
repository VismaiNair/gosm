/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gosm",
	Short: "A CLI tool that simplifies building Go applications for WebAssembly. Eliminates the need to remember complex build flags and streamlines the Go-to-WASM workflow.",
	Long: `gosm is a command-line interface tool built with Cobra that streamlines the process of compiling Go applications to WebAssembly (WASM). 
	Instead of manually typing GOOS=js GOARCH=wasm go build every time, developers can simply use gosm build to compile their Go code for web deployment.

WebAssembly represents the future of web development by enabling near-native performance in browsers while allowing developers to use languages beyond JavaScript. 
However, the current toolchain requires remembering specific environment variables and build commands that can slow down development workflows.

gosm addresses this friction by providing an intuitive CLI that abstracts away the complexity of cross-compilation flags. 

The tool is designed to be the foundation for a comprehensive Go WebAssembly development toolkit.
Built for developers who want to leverage Go's simplicity and performance in web applications without the overhead of complex build configurations. Whether you're building interactive web applications, browser-based tools, or experimenting with WebAssembly's capabilities, gosm makes the Go-to-WASM workflow as simple as gosm build.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gosm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

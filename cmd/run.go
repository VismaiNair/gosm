package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/browser"

	"github.com/spf13/cobra"
)

var port int

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "gosm run serves all static files over a webserver.",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		open, _ := cmd.Flags().GetBool("open")

		fs := http.FileServer(http.Dir("."))
		http.Handle("/", fs)
		fmt.Fprintf(os.Stdout, "The WASM static webserver is starting on http://localhost:%d\n", port)
		fmt.Fprintf(os.Stdout, "Press Ctrl+C or Command+C to stop the server.")
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			return fmt.Errorf("failed to start server: %w", err)
		}
		if open {
			browser.OpenURL(fmt.Sprintf("http://localhost:%d", port))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().IntVarP(&port, "port", "p", 8080, "The port to run the web server on")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

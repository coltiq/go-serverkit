/*
Copyright © 2024 Coltiq <dev@coltiq.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-serverkit",
	Short: "A program to start up an opinionated Go web server",
	Long: `Go ServerKit is a CLI tool designed to streamline the process of setting up a Go web server. 
With Go ServerKit, you can quickly initialize a fully functional web server with a robust structure, complete with 
essential features like database implementation, logging, error handling, and basic routing. This tool is perfect for 
developers who want to expedite their workflow and focus on building web applications without the hassle of setting up 
the foundational components from scratch. Whether you're starting a new project or need a reliable base for rapid 
development, Go ServerKit provides everything you need to get your server up and running efficiently.`,
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
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

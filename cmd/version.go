/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display Go ServerKit version information",
	Long:  `The version command provides information about Go ServerKit's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		version := getGoServerKitVersion()
		fmt.Printf("Go ServerKit CLI version: %s\n", version)
	},
}

func getGoServerKitVersion() string {
	noVersionAvailable := `This build has no version information available.
Run 'go-serverkit help version' for additional information`
	return noVersionAvailable
}

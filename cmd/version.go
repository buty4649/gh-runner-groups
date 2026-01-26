package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is the current version of gh-runner-groups
var Version = "1.0.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gh-runner-groups",
	Long:  `Print the version number of gh-runner-groups`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
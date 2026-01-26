package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-runner-group",
	Short: "A CLI tool for managing GitHub Actions runner groups",
	Long: `A command line tool for interacting with GitHub runner groups
and retrieving runner information in both Enterprise and Organization contexts.

This tool allows you to:
- List runner groups in enterprises or organizations
- List runners in specific runner groups with status information
- Format the output for further processing

Supports both GitHub.com and GitHub Enterprise Server.`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-runner-group.yaml)")
}

// addSubcommands adds all subcommands to the root command
func addSubcommands() {
	rootCmd.AddCommand(runnersCmd)
	rootCmd.AddCommand(listCmd)
}

func init() {
	addSubcommands()
}
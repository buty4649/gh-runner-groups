package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/buty4649/gh-runner-groups/pkg/runnergroup"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List runner groups in an enterprise",
	Long: `List all runner groups in the specified enterprise.

The command requires:
- An enterprise name specified with the --enterprise flag

Optional:
- A hostname specified with the --hostname flag for GitHub Enterprise Server
- The GH_HOST environment variable is also supported (handled by gh CLI)

Examples:
  # For GitHub.com
  gh-runner-group list --enterprise myorg

  # For GitHub Enterprise Server (using flag)
  gh-runner-group list --enterprise myorg --hostname github.example.com

  # For GitHub Enterprise Server (using environment variable)
  GH_HOST=github.example.com gh-runner-group list --enterprise myorg`,
	Args: cobra.NoArgs,
	Run:  runListCommand,
}

func init() {
	// Add the --enterprise flag (shared with runners command)
	listCmd.Flags().StringVarP(&enterpriseName, "enterprise", "e", "", "Enterprise name (required)")
	listCmd.MarkFlagRequired("enterprise")

	// Add the --hostname flag (shared with runners command)
	listCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "GitHub hostname (e.g., github.example.com)")
}

func runListCommand(cmd *cobra.Command, args []string) {
	// Create API client with optional hostname
	// Only pass hostname if explicitly provided via flag (gh handles GH_HOST env var automatically)
	client := runnergroup.NewClient()
	if hostname != "" {
		client.WithHostname(hostname)
	}

	// Get runner groups using the client
	runnerGroups, err := client.ListRunnerGroups(enterpriseName)
	if err != nil {
		log.Fatal(err)
	}

	// Format and print runner groups
	output := runnergroup.FormatRunnerGroups(runnerGroups)
	fmt.Println(output)
}
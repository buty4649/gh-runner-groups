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
	Short: "List runner groups in an enterprise or organization",
	Long: `List all runner groups in the specified enterprise or organization.

The command requires exactly one of:
- An enterprise name specified with the --enterprise flag
- An organization name specified with the --org flag

Optional:
- A hostname specified with the --hostname flag for GitHub Enterprise Server
- The GH_HOST environment variable is also supported (handled by gh CLI)

Examples:
  # For GitHub.com enterprise
  gh-runner-group list --enterprise myenterprise

  # For GitHub.com organization
  gh-runner-group list --org myorg

  # For GitHub Enterprise Server (using flag)
  gh-runner-group list --enterprise myenterprise --hostname github.example.com
  gh-runner-group list --org myorg --hostname github.example.com

  # For GitHub Enterprise Server (using environment variable)
  GH_HOST=github.example.com gh-runner-group list --enterprise myenterprise
  GH_HOST=github.example.com gh-runner-group list --org myorg`,
	Args: cobra.NoArgs,
	Run:  runListCommand,
}

func init() {
	// Add the --enterprise flag (shared with runners command)
	listCmd.Flags().StringVarP(&enterpriseName, "enterprise", "e", "", "Enterprise name")

	// Add the --org flag (shared with runners command)
	listCmd.Flags().StringVarP(&orgName, "org", "o", "", "Organization name")

	// Add the --hostname flag (shared with runners command)
	listCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "GitHub hostname (e.g., github.example.com)")

	// Make enterprise and org mutually exclusive, at least one is required
	listCmd.MarkFlagsMutuallyExclusive("enterprise", "org")
	listCmd.MarkFlagsOneRequired("enterprise", "org")
}

func runListCommand(cmd *cobra.Command, args []string) {
	// Create API client with optional hostname
	// Only pass hostname if explicitly provided via flag (gh handles GH_HOST env var automatically)
	client := runnergroup.NewClient()
	if hostname != "" {
		client.WithHostname(hostname)
	}

	// Get runner groups using the client - choose enterprise or org API based on flags
	var runnerGroups []runnergroup.RunnerGroup
	var err error

	if enterpriseName != "" {
		runnerGroups, err = client.ListRunnerGroups(enterpriseName)
	} else {
		runnerGroups, err = client.ListOrgRunnerGroups(orgName)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Format and print runner groups
	output := runnergroup.FormatRunnerGroups(runnerGroups)
	fmt.Println(output)
}
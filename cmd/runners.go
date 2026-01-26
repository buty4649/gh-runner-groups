package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/buty4649/gh-runner-groups/pkg/runnergroup"
)

// runnersCmd represents the runners command
var runnersCmd = &cobra.Command{
	Use:   "runners <runner-group-id>",
	Short: "List runners in a specific runner group",
	Long: `List all runners in the specified runner group.

The command requires:
- An enterprise name specified with the --enterprise flag
- A runner group ID as a positional argument

Optional:
- A hostname specified with the --hostname flag for GitHub Enterprise Server
- The GH_HOST environment variable is also supported (handled by gh CLI)

Examples:
  # For GitHub.com
  gh-runner-group runners 123 --enterprise myorg

  # For GitHub Enterprise Server (using flag)
  gh-runner-group runners 123 --enterprise myorg --hostname github.example.com

  # For GitHub Enterprise Server (using environment variable)
  GH_HOST=github.example.com gh-runner-group runners 123 --enterprise myorg`,
	Args: cobra.ExactArgs(1),
	Run:  runRunnersCommand,
}

var (
	enterpriseName string
	hostname       string
)

func init() {
	// Add the --enterprise flag
	runnersCmd.Flags().StringVarP(&enterpriseName, "enterprise", "e", "", "Enterprise name (required)")
	if err := runnersCmd.MarkFlagRequired("enterprise"); err != nil {
		log.Fatal(err)
	}

	// Add the --hostname flag
	runnersCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "GitHub hostname (e.g., github.example.com)")
}

func runRunnersCommand(cmd *cobra.Command, args []string) {
	runnerGroupID := args[0]

	// Create API client with optional hostname
	// Only pass hostname if explicitly provided via flag (gh handles GH_HOST env var automatically)
	client := runnergroup.NewClient()
	if hostname != "" {
		client.WithHostname(hostname)
	}

	// Get runners using the client
	runners, err := client.GetRunners(enterpriseName, runnerGroupID)
	if err != nil {
		log.Fatal(err)
	}

	// Sort runners by status (Active -> Idle -> Offline) then by name
	runnergroup.SortRunners(runners)

	// Calculate max name width for alignment
	nameWidth := runnergroup.GetMaxRunnerNameLength(runners)

	// Print aligned header
	runnergroup.PrintHeaderAligned(nameWidth)

	// Output with aligned colored status
	for _, r := range runners {
		fmt.Println(runnergroup.FormatRunnerWithStatusAligned(r, nameWidth))
	}
}
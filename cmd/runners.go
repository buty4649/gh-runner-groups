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
- Exactly one of: an enterprise name (--enterprise flag) OR an organization name (--org flag)
- A runner group ID as a positional argument

Optional:
- A hostname specified with the --hostname flag for GitHub Enterprise Server
- The GH_HOST environment variable is also supported (handled by gh CLI)

Examples:
  # For GitHub.com enterprise
  gh-runner-group runners 123 --enterprise myenterprise

  # For GitHub.com organization
  gh-runner-group runners 123 --org myorg

  # For GitHub Enterprise Server (using flag)
  gh-runner-group runners 123 --enterprise myenterprise --hostname github.example.com
  gh-runner-group runners 123 --org myorg --hostname github.example.com

  # For GitHub Enterprise Server (using environment variable)
  GH_HOST=github.example.com gh-runner-group runners 123 --enterprise myenterprise
  GH_HOST=github.example.com gh-runner-group runners 123 --org myorg`,
	Args: cobra.ExactArgs(1),
	Run:  runRunnersCommand,
}

var (
	enterpriseName string
	orgName        string
	hostname       string
)

func init() {
	// Add the --enterprise flag
	runnersCmd.Flags().StringVarP(&enterpriseName, "enterprise", "e", "", "Enterprise name")

	// Add the --org flag
	runnersCmd.Flags().StringVarP(&orgName, "org", "o", "", "Organization name")

	// Add the --hostname flag
	runnersCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "GitHub hostname (e.g., github.example.com)")

	// Make enterprise and org mutually exclusive, at least one is required
	runnersCmd.MarkFlagsMutuallyExclusive("enterprise", "org")
	runnersCmd.MarkFlagsOneRequired("enterprise", "org")
}

func runRunnersCommand(cmd *cobra.Command, args []string) {
	runnerGroupID := args[0]

	// Create API client with optional hostname
	// Only pass hostname if explicitly provided via flag (gh handles GH_HOST env var automatically)
	client := runnergroup.NewClient()
	if hostname != "" {
		client.WithHostname(hostname)
	}

	// Get runners using the client - choose enterprise or org API based on flags
	var runners []runnergroup.Runner
	var err error

	if enterpriseName != "" {
		runners, err = client.GetRunners(enterpriseName, runnerGroupID)
	} else {
		runners, err = client.GetOrgRunners(orgName, runnerGroupID)
	}

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
package runnergroup

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// GetRunners fetches runners from the specified enterprise and runner group
func (c *Client) GetRunners(enterpriseID, runnerGroupID string) ([]Runner, error) {
	// Validate runner group ID is a number
	if _, err := strconv.Atoi(runnerGroupID); err != nil {
		return nil, fmt.Errorf("invalid runner group ID: %s (must be a number)", runnerGroupID)
	}

	var allRunners []Runner
	page := 1
	perPage := 100

	for {
		// Build API endpoint with pagination parameters
		endpoint := fmt.Sprintf("/enterprises/%s/actions/runner-groups/%s/runners?per_page=%d&page=%d",
			enterpriseID, runnerGroupID, perPage, page)

		// Call GitHub API
		var response RunnersResponse
		if err := c.CallAPIWithJSON(endpoint, &response); err != nil {
			return nil, err
		}

		// Add runners from this page
		allRunners = append(allRunners, response.Runners...)

		// If we got fewer runners than per_page, this was the last page
		if len(response.Runners) < perPage {
			break
		}

		page++
	}

	return allRunners, nil
}


// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorOrange = "\033[33m"
	ColorGray   = "\033[37m"
)


// getStatusPriority returns priority for sorting (lower number = higher priority)
func getStatusPriority(runner Runner) int {
	if runner.Status == "online" && runner.Busy {
		return 0 // Active - highest priority
	} else if runner.Status == "online" && !runner.Busy {
		return 1 // Idle - medium priority
	} else {
		return 2 // Offline - lowest priority
	}
}

// SortRunners sorts runners by status (Active -> Idle -> Offline) then by name
func SortRunners(runners []Runner) {
	sort.Slice(runners, func(i, j int) bool {
		// First sort by status priority
		priorityI := getStatusPriority(runners[i])
		priorityJ := getStatusPriority(runners[j])

		if priorityI != priorityJ {
			return priorityI < priorityJ
		}

		// If same status, sort by runner name
		return runners[i].Name < runners[j].Name
	})
}

// GetMaxRunnerNameLength returns the maximum length of runner names
func GetMaxRunnerNameLength(runners []Runner) int {
	maxLen := len("Runners") // Header length as minimum
	for _, runner := range runners {
		if len(runner.Name) > maxLen {
			maxLen = len(runner.Name)
		}
	}
	return maxLen
}

// FormatRunnerWithStatus formats a runner with colored status
func FormatRunnerWithStatus(runner Runner) string {
	var status string

	if runner.Status == "online" && runner.Busy {
		// Active (orange)
		status = fmt.Sprintf("%s● Active%s", ColorOrange, ColorReset)
	} else if runner.Status == "online" && !runner.Busy {
		// Idle (green)
		status = fmt.Sprintf("%s● Idle%s", ColorGreen, ColorReset)
	} else {
		// Offline (gray)
		status = fmt.Sprintf("%s● Offline%s", ColorGray, ColorReset)
	}

	return fmt.Sprintf("%s\t%s", runner.Name, status)
}

// FormatRunnerWithStatusAligned formats a runner with colored status and aligned columns
func FormatRunnerWithStatusAligned(runner Runner, nameWidth int) string {
	var status string

	if runner.Status == "online" && runner.Busy {
		// Active (orange)
		status = fmt.Sprintf("%s● Active%s", ColorOrange, ColorReset)
	} else if runner.Status == "online" && !runner.Busy {
		// Idle (green)
		status = fmt.Sprintf("%s● Idle%s", ColorGreen, ColorReset)
	} else {
		// Offline (gray)
		status = fmt.Sprintf("%s● Offline%s", ColorGray, ColorReset)
	}

	// Pad runner name to align columns
	paddedName := runner.Name + strings.Repeat(" ", nameWidth-len(runner.Name))
	return fmt.Sprintf("%s  %s", paddedName, status)
}


// PrintHeaderAligned prints the table header with aligned columns
func PrintHeaderAligned(nameWidth int) {
	paddedHeader := "Runners" + strings.Repeat(" ", nameWidth-len("Runners"))
	fmt.Printf("%s  %s\n", paddedHeader, "Status")
}


// GetMaxRunnerGroupNameLength returns the maximum length of runner group names
func GetMaxRunnerGroupNameLength(groups []RunnerGroup) int {
	maxLen := len("Name") // Header length as minimum
	for _, group := range groups {
		if len(group.Name) > maxLen {
			maxLen = len(group.Name)
		}
	}
	return maxLen
}

// FormatRunnerGroupWithStatus formats a runner group with visibility info
func FormatRunnerGroupWithStatus(group RunnerGroup, nameWidth int) string {
	var status string

	if group.Default {
		// Default group (green)
		status = fmt.Sprintf("%s● %s (default)%s", ColorGreen, group.Visibility, ColorReset)
	} else if group.Visibility == "private" {
		// Private (gray)
		status = fmt.Sprintf("%s● %s%s", ColorGray, group.Visibility, ColorReset)
	} else {
		// Selected/All (orange)
		status = fmt.Sprintf("%s● %s%s", ColorOrange, group.Visibility, ColorReset)
	}

	// Pad group name to align columns
	paddedName := group.Name + strings.Repeat(" ", nameWidth-len(group.Name))
	return fmt.Sprintf("%d\t%s  %s", group.ID, paddedName, status)
}

// PrintRunnerGroupHeaderAligned prints the runner groups table header with aligned columns
func PrintRunnerGroupHeaderAligned(nameWidth int) {
	paddedHeader := "Name" + strings.Repeat(" ", nameWidth-len("Name"))
	fmt.Printf("ID\t%s  Visibility\n", paddedHeader)
}
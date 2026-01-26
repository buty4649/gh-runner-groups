package runnergroup

import (
	"fmt"
	"strings"
)

// ListRunnerGroups fetches runner groups from the specified enterprise
func (c *Client) ListRunnerGroups(enterpriseID string) ([]RunnerGroup, error) {
	var allRunnerGroups []RunnerGroup
	page := 1
	perPage := 100

	for {
		// Build API endpoint with pagination parameters
		endpoint := fmt.Sprintf("/enterprises/%s/actions/runner-groups?per_page=%d&page=%d",
			enterpriseID, perPage, page)

		// Call GitHub API
		var response RunnerGroupsResponse
		if err := c.CallAPIWithJSON(endpoint, &response); err != nil {
			return nil, err
		}

		// Add runner groups from this page
		allRunnerGroups = append(allRunnerGroups, response.RunnerGroups...)

		// If we got fewer runner groups than per_page, this was the last page
		if len(response.RunnerGroups) < perPage {
			break
		}

		page++
	}

	return allRunnerGroups, nil
}

// ListOrgRunnerGroups fetches runner groups from the specified organization
func (c *Client) ListOrgRunnerGroups(org string) ([]RunnerGroup, error) {
	var allRunnerGroups []RunnerGroup
	page := 1
	perPage := 100

	for {
		// Build API endpoint with pagination parameters
		endpoint := fmt.Sprintf("/orgs/%s/actions/runner-groups?per_page=%d&page=%d",
			org, perPage, page)

		// Call GitHub API
		var response RunnerGroupsResponse
		if err := c.CallAPIWithJSON(endpoint, &response); err != nil {
			return nil, err
		}

		// Add runner groups from this page
		allRunnerGroups = append(allRunnerGroups, response.RunnerGroups...)

		// If we got fewer runner groups than per_page, this was the last page
		if len(response.RunnerGroups) < perPage {
			break
		}

		page++
	}

	return allRunnerGroups, nil
}

// FormatRunnerGroups formats runner groups for display
func FormatRunnerGroups(groups []RunnerGroup) string {
	if len(groups) == 0 {
		return ""
	}

	// Calculate max name width for alignment
	nameWidth := GetMaxRunnerGroupNameLength(groups)

	// Create header
	paddedHeader := "Name" + strings.Repeat(" ", nameWidth-len("Name"))
	header := fmt.Sprintf("ID\t%s  Visibility", paddedHeader)

	// Create formatted lines
	lines := []string{header}
	for _, group := range groups {
		line := FormatRunnerGroupWithStatus(group, nameWidth)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
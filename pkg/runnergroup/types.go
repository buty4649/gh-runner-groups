package runnergroup

// Runner represents a GitHub Actions runner
type Runner struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Busy   bool   `json:"busy"`
}

// RunnersResponse represents the API response containing runners
type RunnersResponse struct {
	Runners []Runner `json:"runners"`
}

// RunnerGroup represents a GitHub Actions runner group
type RunnerGroup struct {
	ID                         int      `json:"id"`
	Name                       string   `json:"name"`
	Visibility                 string   `json:"visibility"`
	Default                    bool     `json:"default"`
	SelectedRepositoriesURL    string   `json:"selected_repositories_url"`
	RunnersURL                 string   `json:"runners_url"`
	Inherited                  bool     `json:"inherited"`
	AllowsPublicRepositories   bool     `json:"allows_public_repositories"`
	RestrictedToWorkflows      bool     `json:"restricted_to_workflows"`
	SelectedWorkflows          []string `json:"selected_workflows"`
}

// RunnerGroupsResponse represents the API response containing runner groups
type RunnerGroupsResponse struct {
	TotalCount   int           `json:"total_count"`
	RunnerGroups []RunnerGroup `json:"runner_groups"`
}

// Options represents options for GitHub API calls
type Options struct {
	Headers   map[string]string
	Paginate  bool
	Hostname  string
}
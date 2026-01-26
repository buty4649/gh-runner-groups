package runnergroup

import (
	"encoding/json"
	"fmt"

	"github.com/cli/go-gh/v2"
)

// Client represents a GitHub API client with configured options
type Client struct {
	Options Options
}

// NewClient creates a new GitHub API client with default options
func NewClient() *Client {
	return &Client{
		Options: Options{
			Headers: map[string]string{
				"Accept":               "application/vnd.github+json",
				"X-GitHub-Api-Version": "2022-11-28",
			},
			Paginate: false, // Disable automatic pagination, use manual pagination instead
		},
	}
}

// WithHostname sets the hostname for the client
func (c *Client) WithHostname(hostname string) *Client {
	c.Options.Hostname = hostname
	return c
}

// CallAPI makes a GitHub API call and returns the raw response using the client's options
func (c *Client) CallAPI(endpoint string) ([]byte, error) {
	args := []string{"api"}

	// Add headers
	for key, value := range c.Options.Headers {
		args = append(args, "-H", fmt.Sprintf("%s: %s", key, value))
	}

	// Add hostname if specified
	if c.Options.Hostname != "" {
		args = append(args, "--hostname", c.Options.Hostname)
	}

	// Note: Manual pagination is handled in the caller, not here

	// Add endpoint
	args = append(args, endpoint)

	stdout, stderr, err := gh.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute gh command: %v\nStderr: %s", err, stderr.String())
	}

	return stdout.Bytes(), nil
}

// CallAPIWithJSON makes a GitHub API call and unmarshals the JSON response using the client's options
func (c *Client) CallAPIWithJSON(endpoint string, result interface{}) error {
	data, err := c.CallAPI(endpoint)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return nil
}
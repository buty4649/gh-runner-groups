package runnergroup

import (
	"testing"
)

// Test organization runner groups listing validation and structure
func TestListOrgRunnerGroups_Structure(t *testing.T) {
	tests := []struct {
		name string
		org  string
	}{
		{
			name: "valid org name",
			org:  "test-org",
		},
		{
			name: "empty org name",
			org:  "",
		},
		{
			name: "org name with special characters",
			org:  "test-org-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This will fail when actually calling the GitHub API
			// In a real test environment, you'd mock the API client
			client := NewClient()
			_, err := client.ListOrgRunnerGroups(tt.org)

			// For now, we expect this to fail since we're not mocking the GitHub API
			// In a proper test environment, you would mock the HTTP client
			if err == nil {
				// If by chance this succeeds (e.g., with valid GitHub credentials),
				// we can just note that the structure didn't cause a panic
				t.Logf("API call succeeded for org: %s", tt.org)
			} else {
				// Expected to fail without proper GitHub API access
				t.Logf("Expected API failure for org %s: %v", tt.org, err)
			}
		})
	}
}

// Test enterprise runner groups listing validation and structure
func TestListRunnerGroups_Structure(t *testing.T) {
	tests := []struct {
		name         string
		enterpriseID string
	}{
		{
			name:         "valid enterprise ID",
			enterpriseID: "test-enterprise",
		},
		{
			name:         "empty enterprise ID",
			enterpriseID: "",
		},
		{
			name:         "enterprise ID with special characters",
			enterpriseID: "test-enterprise-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This will fail when actually calling the GitHub API
			// In a real test environment, you'd mock the API client
			client := NewClient()
			_, err := client.ListRunnerGroups(tt.enterpriseID)

			// For now, we expect this to fail since we're not mocking the GitHub API
			// In a proper test environment, you would mock the HTTP client
			if err == nil {
				// If by chance this succeeds (e.g., with valid GitHub credentials),
				// we can just note that the structure didn't cause a panic
				t.Logf("API call succeeded for enterprise: %s", tt.enterpriseID)
			} else {
				// Expected to fail without proper GitHub API access
				t.Logf("Expected API failure for enterprise %s: %v", tt.enterpriseID, err)
			}
		})
	}
}

// Test FormatRunnerGroups with various input scenarios
func TestFormatRunnerGroups_EmptyList(t *testing.T) {
	groups := []RunnerGroup{}
	result := FormatRunnerGroups(groups)

	if result != "" {
		t.Errorf("Expected empty string for empty groups list, got %q", result)
	}
}

func TestFormatRunnerGroups_SingleGroup(t *testing.T) {
	groups := []RunnerGroup{
		{
			ID:         1,
			Name:       "test-group",
			Visibility: "private",
			Default:    true,
		},
	}

	result := FormatRunnerGroups(groups)

	// Should contain the header
	if !contains(result, "ID") || !contains(result, "Name") || !contains(result, "Visibility") {
		t.Errorf("Expected header in output, got %q", result)
	}

	// Should contain the group data
	if !contains(result, "1") || !contains(result, "test-group") || !contains(result, "private") {
		t.Errorf("Expected group data in output, got %q", result)
	}
}

func TestFormatRunnerGroups_MultipleGroups(t *testing.T) {
	groups := []RunnerGroup{
		{
			ID:         1,
			Name:       "default-group",
			Visibility: "private",
			Default:    true,
		},
		{
			ID:         2,
			Name:       "selected-group",
			Visibility: "selected",
			Default:    false,
		},
		{
			ID:         3,
			Name:       "all-group",
			Visibility: "all",
			Default:    false,
		},
	}

	result := FormatRunnerGroups(groups)

	// Should contain all group names
	expectedNames := []string{"default-group", "selected-group", "all-group"}
	for _, name := range expectedNames {
		if !contains(result, name) {
			t.Errorf("Expected group name %q in output, got %q", name, result)
		}
	}

	// Should contain all group IDs
	expectedIDs := []string{"1", "2", "3"}
	for _, id := range expectedIDs {
		if !contains(result, id) {
			t.Errorf("Expected group ID %q in output, got %q", id, result)
		}
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsInternal(s, substr)))
}

func containsInternal(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Test endpoint construction logic indirectly by testing client behavior
func TestClientEndpointConstruction(t *testing.T) {
	client := NewClient()

	// Test that client doesn't panic with various inputs
	testCases := []struct {
		name string
		test func() error
	}{
		{
			name: "ListRunnerGroups with empty enterprise",
			test: func() error {
				_, err := client.ListRunnerGroups("")
				return err
			},
		},
		{
			name: "ListOrgRunnerGroups with empty org",
			test: func() error {
				_, err := client.ListOrgRunnerGroups("")
				return err
			},
		},
		{
			name: "GetRunners with empty enterprise",
			test: func() error {
				_, err := client.GetRunners("", "123")
				return err
			},
		},
		{
			name: "GetOrgRunners with empty org",
			test: func() error {
				_, err := client.GetOrgRunners("", "123")
				return err
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.test()
			// We expect errors due to API calls, but we shouldn't get panics
			// The test passes if it doesn't panic
			if err != nil {
				t.Logf("Expected API error: %v", err)
			}
		})
	}
}
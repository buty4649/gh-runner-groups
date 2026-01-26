package runnergroup

import (
	"strings"
	"testing"
)


func TestGetStatusPriority(t *testing.T) {
	tests := []struct {
		name     string
		runner   Runner
		expected int
	}{
		{
			name: "active runner (highest priority)",
			runner: Runner{
				Status: "online",
				Busy:   true,
			},
			expected: 0,
		},
		{
			name: "idle runner (medium priority)",
			runner: Runner{
				Status: "online",
				Busy:   false,
			},
			expected: 1,
		},
		{
			name: "offline runner (lowest priority)",
			runner: Runner{
				Status: "offline",
				Busy:   false,
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getStatusPriority(tt.runner)
			if result != tt.expected {
				t.Errorf("Expected priority %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestSortRunners(t *testing.T) {
	runners := []Runner{
		{Name: "z-offline", Status: "offline", Busy: false},
		{Name: "a-idle", Status: "online", Busy: false},
		{Name: "z-active", Status: "online", Busy: true},
		{Name: "a-offline", Status: "offline", Busy: false},
		{Name: "z-idle", Status: "online", Busy: false},
		{Name: "a-active", Status: "online", Busy: true},
	}

	SortRunners(runners)

	expectedOrder := []string{
		"a-active", // Active runners first, sorted by name
		"z-active",
		"a-idle", // Idle runners second, sorted by name
		"z-idle",
		"a-offline", // Offline runners last, sorted by name
		"z-offline",
	}

	for i, runner := range runners {
		if runner.Name != expectedOrder[i] {
			t.Errorf("Expected runner at position %d to be %s, got %s", i, expectedOrder[i], runner.Name)
		}
	}
}

func TestGetMaxRunnerNameLength(t *testing.T) {
	tests := []struct {
		name     string
		runners  []Runner
		expected int
	}{
		{
			name:     "empty runners list",
			runners:  []Runner{},
			expected: len("Runners"), // Header length as minimum
		},
		{
			name: "runners with various name lengths",
			runners: []Runner{
				{Name: "short"},
				{Name: "very-long-runner-name"},
				{Name: "mid"},
			},
			expected: len("very-long-runner-name"),
		},
		{
			name: "all names shorter than header",
			runners: []Runner{
				{Name: "a"},
				{Name: "bb"},
				{Name: "ccc"},
			},
			expected: len("Runners"), // Header length as minimum
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetMaxRunnerNameLength(tt.runners)
			if result != tt.expected {
				t.Errorf("Expected max length %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestFormatRunnerWithStatus(t *testing.T) {
	tests := []struct {
		name     string
		runner   Runner
		contains []string // Strings that should be in the output
	}{
		{
			name: "active runner",
			runner: Runner{
				Name:   "test-runner",
				Status: "online",
				Busy:   true,
			},
			contains: []string{"test-runner", "● Active"},
		},
		{
			name: "idle runner",
			runner: Runner{
				Name:   "idle-runner",
				Status: "online",
				Busy:   false,
			},
			contains: []string{"idle-runner", "● Idle"},
		},
		{
			name: "offline runner",
			runner: Runner{
				Name:   "offline-runner",
				Status: "offline",
				Busy:   false,
			},
			contains: []string{"offline-runner", "● Offline"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatRunnerWithStatus(tt.runner)
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, got %q", expected, result)
				}
			}
		})
	}
}

func TestFormatRunnerWithStatusAligned(t *testing.T) {
	runner := Runner{
		Name:   "test",
		Status: "online",
		Busy:   true,
	}

	nameWidth := 10
	result := FormatRunnerWithStatusAligned(runner, nameWidth)

	// Should contain the runner name
	if !strings.Contains(result, "test") {
		t.Errorf("Expected output to contain runner name, got %q", result)
	}

	// Should contain status
	if !strings.Contains(result, "● Active") {
		t.Errorf("Expected output to contain status, got %q", result)
	}

	// Should have proper alignment (runner name + spaces + status)
	parts := strings.Split(result, "  ") // Split on double space
	if len(parts) < 2 {
		t.Errorf("Expected aligned output with double space separator, got %q", result)
	}
}


func TestGetMaxRunnerGroupNameLength(t *testing.T) {
	tests := []struct {
		name     string
		groups   []RunnerGroup
		expected int
	}{
		{
			name:     "empty groups list",
			groups:   []RunnerGroup{},
			expected: len("Name"), // Header length as minimum
		},
		{
			name: "groups with various name lengths",
			groups: []RunnerGroup{
				{Name: "short"},
				{Name: "very-long-group-name"},
				{Name: "mid"},
			},
			expected: len("very-long-group-name"),
		},
		{
			name: "all names shorter than header",
			groups: []RunnerGroup{
				{Name: "a"},
				{Name: "bb"},
			},
			expected: len("Name"), // Header length as minimum
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetMaxRunnerGroupNameLength(tt.groups)
			if result != tt.expected {
				t.Errorf("Expected max length %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestFormatRunnerGroupWithStatus(t *testing.T) {
	tests := []struct {
		name     string
		group    RunnerGroup
		contains []string // Strings that should be in the output
	}{
		{
			name: "default group",
			group: RunnerGroup{
				ID:         1,
				Name:       "default-group",
				Visibility: "private",
				Default:    true,
			},
			contains: []string{"1", "default-group", "● private (default)"},
		},
		{
			name: "private group",
			group: RunnerGroup{
				ID:         2,
				Name:       "private-group",
				Visibility: "private",
				Default:    false,
			},
			contains: []string{"2", "private-group", "● private"},
		},
		{
			name: "selected group",
			group: RunnerGroup{
				ID:         3,
				Name:       "selected-group",
				Visibility: "selected",
				Default:    false,
			},
			contains: []string{"3", "selected-group", "● selected"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nameWidth := 20
			result := FormatRunnerGroupWithStatus(tt.group, nameWidth)
			for _, expected := range tt.contains {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected output to contain %q, got %q", expected, result)
				}
			}
		})
	}
}

func TestGetRunners_ValidationError(t *testing.T) {
	tests := []struct {
		name          string
		enterpriseID  string
		runnerGroupID string
		expectError   bool
		expectedError string
	}{
		{
			name:          "valid numeric runner group ID",
			enterpriseID:  "test-enterprise",
			runnerGroupID: "123",
			expectError:   false,
		},
		{
			name:          "invalid non-numeric runner group ID",
			enterpriseID:  "test-enterprise",
			runnerGroupID: "abc",
			expectError:   true,
			expectedError: "invalid runner group ID: abc (must be a number)",
		},
		{
			name:          "empty runner group ID",
			enterpriseID:  "test-enterprise",
			runnerGroupID: "",
			expectError:   true,
			expectedError: "invalid runner group ID:  (must be a number)",
		},
		{
			name:          "negative runner group ID",
			enterpriseID:  "test-enterprise",
			runnerGroupID: "-1",
			expectError:   false, // strconv.Atoi accepts negative numbers
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This will fail when actually calling the GitHub API
			// In a real test environment, you'd mock the API client
			client := NewClient()
			_, err := client.GetRunners(tt.enterpriseID, tt.runnerGroupID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
					return
				}
				if err.Error() != tt.expectedError {
					t.Errorf("Expected error %q, got %q", tt.expectedError, err.Error())
				}
			}
		})
	}
}

func TestGetOrgRunners_ValidationError(t *testing.T) {
	tests := []struct {
		name          string
		org           string
		runnerGroupID string
		expectError   bool
		expectedError string
	}{
		{
			name:          "valid numeric runner group ID",
			org:           "test-org",
			runnerGroupID: "123",
			expectError:   false,
		},
		{
			name:          "invalid non-numeric runner group ID",
			org:           "test-org",
			runnerGroupID: "abc",
			expectError:   true,
			expectedError: "invalid runner group ID: abc (must be a number)",
		},
		{
			name:          "empty runner group ID",
			org:           "test-org",
			runnerGroupID: "",
			expectError:   true,
			expectedError: "invalid runner group ID:  (must be a number)",
		},
		{
			name:          "negative runner group ID",
			org:           "test-org",
			runnerGroupID: "-1",
			expectError:   false, // strconv.Atoi accepts negative numbers
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: This will fail when actually calling the GitHub API
			// In a real test environment, you'd mock the API client
			client := NewClient()
			_, err := client.GetOrgRunners(tt.org, tt.runnerGroupID)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
					return
				}
				if err.Error() != tt.expectedError {
					t.Errorf("Expected error %q, got %q", tt.expectedError, err.Error())
				}
			}
		})
	}
}
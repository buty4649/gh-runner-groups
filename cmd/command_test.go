package cmd

import (
	"strings"
	"testing"
)

// Test help text content for list command
func TestListCommand_HelpContent(t *testing.T) {
	help := listCmd.Long

	expectedStrings := []string{
		"enterprise or organization",
		"exactly one of",
		"--enterprise flag",
		"--org flag",
		"GitHub.com enterprise",
		"GitHub.com organization",
		"GitHub Enterprise Server",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(help, expected) {
			t.Errorf("Expected help text to contain %q, but it doesn't. Help text: %q", expected, help)
		}
	}
}

// Test help text content for runners command
func TestRunnersCommand_HelpContent(t *testing.T) {
	help := runnersCmd.Long

	expectedStrings := []string{
		"Exactly one of",
		"enterprise name (--enterprise flag)",
		"organization name (--org flag)",
		"runner group ID as a positional argument",
		"GitHub.com enterprise",
		"GitHub.com organization",
		"GitHub Enterprise Server",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(help, expected) {
			t.Errorf("Expected help text to contain %q, but it doesn't. Help text: %q", expected, help)
		}
	}
}

// Test usage examples in help text for list command
func TestListCommand_UsageExamples(t *testing.T) {
	help := listCmd.Long

	expectedExamples := []string{
		"gh-runner-group list --enterprise myenterprise",
		"gh-runner-group list --org myorg",
		"--hostname github.example.com",
		"GH_HOST=github.example.com",
	}

	for _, example := range expectedExamples {
		if !strings.Contains(help, example) {
			t.Errorf("Expected help text to contain example %q, but it doesn't", example)
		}
	}
}

// Test usage examples in help text for runners command
func TestRunnersCommand_UsageExamples(t *testing.T) {
	help := runnersCmd.Long

	expectedExamples := []string{
		"gh-runner-group runners 123 --enterprise myenterprise",
		"gh-runner-group runners 123 --org myorg",
		"--hostname github.example.com",
		"GH_HOST=github.example.com",
	}

	for _, example := range expectedExamples {
		if !strings.Contains(help, example) {
			t.Errorf("Expected help text to contain example %q, but it doesn't", example)
		}
	}
}

// Test command descriptions
func TestListCommand_Description(t *testing.T) {
	expected := "List runner groups in an enterprise or organization"
	if listCmd.Short != expected {
		t.Errorf("Expected short description %q, got %q", expected, listCmd.Short)
	}
}

func TestRunnersCommand_Description(t *testing.T) {
	expected := "List runners in a specific runner group"
	if runnersCmd.Short != expected {
		t.Errorf("Expected short description %q, got %q", expected, runnersCmd.Short)
	}
}

// Test command usage strings
func TestRunnersCommand_Usage(t *testing.T) {
	expected := "runners <runner-group-id>"
	if runnersCmd.Use != expected {
		t.Errorf("Expected usage %q, got %q", expected, runnersCmd.Use)
	}
}

func TestListCommand_Usage(t *testing.T) {
	expected := "list"
	if listCmd.Use != expected {
		t.Errorf("Expected usage %q, got %q", expected, listCmd.Use)
	}
}

// Test root command description includes organization support
func TestRootCommand_Description(t *testing.T) {
	help := rootCmd.Long

	expectedStrings := []string{
		"Enterprise and Organization contexts",
		"List runner groups in enterprises or organizations",
		"GitHub.com and GitHub Enterprise Server",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(help, expected) {
			t.Errorf("Expected root help text to contain %q, but it doesn't. Help text: %q", expected, help)
		}
	}
}
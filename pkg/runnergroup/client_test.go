package runnergroup

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	expectedHeaders := map[string]string{
		"Accept":               "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
	}

	if !reflect.DeepEqual(client.Options.Headers, expectedHeaders) {
		t.Errorf("Expected headers %v, got %v", expectedHeaders, client.Options.Headers)
	}

	if client.Options.Paginate {
		t.Errorf("Expected Paginate to be false, got %v", client.Options.Paginate)
	}

	if client.Options.Hostname != "" {
		t.Errorf("Expected empty Hostname, got %v", client.Options.Hostname)
	}
}

func TestClient_WithHostname(t *testing.T) {
	client := NewClient()
	client.WithHostname("github.example.com")

	if client.Options.Hostname != "github.example.com" {
		t.Errorf("Expected hostname 'github.example.com', got %v", client.Options.Hostname)
	}
}

// Test JSON unmarshaling functionality with generic interface
func TestJSONUnmarshaling(t *testing.T) {
	// Test basic JSON unmarshaling without depending on runner package

	// Test valid JSON
	validJSON := `{"message": "success", "data": [{"id": 1, "name": "test"}]}`
	var validResult map[string]interface{}
	err := json.Unmarshal([]byte(validJSON), &validResult)

	if err != nil {
		t.Errorf("Expected no error for valid JSON, got %v", err)
	}

	if validResult["message"] != "success" {
		t.Errorf("Expected message 'success', got %v", validResult["message"])
	}

	// Test invalid JSON
	invalidJSON := `{"invalid": json}`
	var invalidResult map[string]interface{}
	err = json.Unmarshal([]byte(invalidJSON), &invalidResult)

	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}

	// Test empty JSON object
	emptyJSON := `{}`
	var emptyResult map[string]interface{}
	err = json.Unmarshal([]byte(emptyJSON), &emptyResult)

	if err != nil {
		t.Errorf("Expected no error for empty JSON object, got %v", err)
	}

	if len(emptyResult) != 0 {
		t.Errorf("Expected empty map, got %v", emptyResult)
	}
}


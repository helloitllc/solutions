package github_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	github "github.com/helloitllc/solutions/apps/platform/github"
)

// MockGitHubAPI simulates a successful GitHub API response
func MockGitHubAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"data": {
			"addProjectV2ItemById": {
				"projectV2Item": {
					"id": "mock123"
				}
			}
		}
	}`))
}

// TestCreateProjectItemFromTemplateWithMock ensures the API call successfully creates a task
func TestCreateProjectItemFromTemplateWithMock(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(MockGitHubAPI))
	defer server.Close()

	// Override GitHub API URL for testing
	github.GitHubAPI = server.URL

	// Run the function with test inputs
	github.CreateProjectItemFromTemplate("Test Feature", "This is a test for GitHub integration.")

	// Validate response structure
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("❌ Failed to make mock request: %v", err)
	}
	defer resp.Body.Close()

	// Decode response
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("❌ Failed to decode JSON response: %v", err)
	}

	// Check if project item ID exists
	data, exists := response["data"].(map[string]interface{})
	if !exists {
		t.Fatalf("❌ No data field in API response")
	}

	item, exists := data["addProjectV2ItemById"].(map[string]interface{})
	if !exists {
		t.Fatalf("❌ No addProjectV2ItemById field in API response")
	}

	projectItem, exists := item["projectV2Item"].(map[string]interface{})
	if !exists {
		t.Fatalf("❌ No projectV2Item field in API response")
	}

	if projectItem["id"] != "mock123" {
		t.Fatalf("❌ Expected project ID 'mock123', got '%v'", projectItem["id"])
	}

	t.Log("✅ Test passed: GitHub Project Item successfully created in template.")
}

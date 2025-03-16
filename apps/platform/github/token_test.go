package github_test

import (
	"os"
	"testing"

	"github.com/helloitllc/solutions/apps/platform/github" // Explicit import
)

// TestGetGitHubToken verifies token retrieval from environment variables
func TestGetGitHubToken(t *testing.T) {
	expectedToken := "mock_github_token"
	os.Setenv("GITHUB_TOKEN", expectedToken) // Mock environment variable

	token := github.GetGitHubToken() // Use full package reference
	if token != expectedToken {
		t.Fatalf("❌ Expected '%s', got '%s'", expectedToken, token)
	}

	t.Log("✅ GitHub Token Retrieved Successfully in Test")
}

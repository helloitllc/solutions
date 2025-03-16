package github_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/helloitllc/solutions/apps/platform/github"
)

// MockGitHubCreateIssueAPI simulates a successful GitHub issue creation response
func MockGitHubCreateIssueAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"id": 789, "title": "Test Issue", "body": "This is a test issue."}`))
}

// TestCreateGitHubIssue verifies that issue creation works correctly
func TestCreateGitHubIssue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(MockGitHubCreateIssueAPI))
	defer server.Close()

	// Override GitHub API Base URL for testing
	github.GitHubAPIBase = server.URL

	success := github.CreateGitHubIssue("mockuser/mockrepo", "Test Issue", "This is a test issue.")
	if !success {
		t.Fatalf("❌ Expected issue creation success but got failure")
	}

	t.Log("✅ GitHub Issue Creation Test Passed")
}

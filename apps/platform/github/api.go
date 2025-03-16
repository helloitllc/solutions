package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GitHub API Base URL (exported for testing)
var GitHubAPIBase = "https://api.github.com"

// GitHub API endpoint for creating issues
const createIssueEndpoint = "/repos/%s/issues"

// GitHubIssueRequest defines the payload for issue creation
type GitHubIssueRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// CreateGitHubIssue creates a new issue in a given GitHub repository
func CreateGitHubIssue(repo, title, body string) bool {
	token := GetGitHubToken()
	if token == "" {
		fmt.Println("❌ No GitHub token found.")
		return false
	}

	// Prepare request payload
	payload := GitHubIssueRequest{Title: title, Body: body}
	payloadBytes, _ := json.Marshal(payload)

	// Format request URL
	apiURL := fmt.Sprintf(GitHubAPIBase+createIssueEndpoint, repo)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("❌ Error creating request:", err)
		return false
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Error making request:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("❌ Failed to create GitHub issue. Status:", resp.Status)
		return false
	}

	fmt.Println("✅ GitHub Issue Created Successfully!")
	return true
}

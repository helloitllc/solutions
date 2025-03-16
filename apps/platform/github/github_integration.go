package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Exported API URL for testing
var GitHubAPI = "https://api.github.com/graphql"

// GitHub request payload
type GitHubRequest struct {
	Query string `json:"query"`
}

// CreateProjectItemFromTemplate adds a new task to the Solution Template project
func CreateProjectItemFromTemplate(title, description string) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		fmt.Println("❌ GITHUB_TOKEN is not set")
		return
	}

	query := fmt.Sprintf(`mutation {
		addProjectV2ItemById(input: {
			projectId: "YOUR_TEMPLATE_PROJECT_ID",
			content: {
				title: "%s",
				body: "%s"
			}
		}) {
			projectV2Item {
				id
			}
		}
	}`, title, description)

	reqBody, _ := json.Marshal(GitHubRequest{Query: query})
	req, _ := http.NewRequest("POST", GitHubAPI, bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer "+githubToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("✅ GitHub Project Item Created - Response Status:", resp.Status)
}

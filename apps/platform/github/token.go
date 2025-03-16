package github

import (
	"fmt"
	"os"
)

// GetGitHubToken retrieves the GitHub token from environment variables
func GetGitHubToken() string {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("❌ GITHUB_TOKEN is not set. Please configure it in your environment.")
		return ""
	}

	fmt.Println("✅ GitHub Token Retrieved Successfully!")
	return token
}

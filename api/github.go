package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/samrat-rm/github_user_activity/models"
)

func FetchGitHubUser(username string) (models.GitHubEvent, error) {
	if len(username) == 0 {
		log.Fatalln("empty username, Please enter a valid username")
	}
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making request: %v\n", err)
		return models.GitHubEvent{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch data: %v\n", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v\n", err)

	}

	var user models.GitHubEvent
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v\n", err)

	}

	return user, nil
}
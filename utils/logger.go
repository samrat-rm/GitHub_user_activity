package utils

import (
	"fmt"

	"github.com/samrat-rm/github_user_activity/models"
)

func GithubActivityLogger(events []models.GitHubEvent) {

	for _, event := range events {

		switch event.Type {
		case "PushEvent":
			fmt.Printf("* Pushed %d commits to %s \n", len(event.Payload.Commits), event.Repo.Name)

		case "CreateEvent":
			fmt.Printf("* Created %s repo \n", event.Repo.Name)

		case "ForkEvent":
			fmt.Printf("* Forked %s repo \n ", event.Repo.Name)

		case "PullRequestEvent":
			fmt.Printf("* PR %sed in %s repo\n ", event.Payload.Action, event.Repo.Name)
		}
	}
}

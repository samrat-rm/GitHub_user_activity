package main

import (
	"os"

	"github.com/samrat-rm/github_user_activity/api"
	"github.com/samrat-rm/github_user_activity/utils"
	"github.com/samrat-rm/github_user_activity/validator"
)

func main() {

	username := validator.InputsValidator(os.Args)
	events := api.FetchGitHubUser(username)

	utils.GithubActivityLogger(events)
}

package validator

import (
	"log"
)

func InputsValidator(args []string) string {
	program := args[0]

	if program != "github-activity" {
		log.Println(program)
		log.Fatalln("invalid program name, Please try using ` \"github-activity\" <username> ` or ` \"./github-activity\" <username> `")
	}

	username := args[1]

	if len(username) == 0 {
		log.Fatalln("empty username, Please enter a valid github username")
	}
	return username
}

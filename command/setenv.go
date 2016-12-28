package command

import (
	"log"
	"github.com/wantedly/developers-account-mapper/store"
	"os"
)

type SetenvCommand struct {
	Meta
}

func (c *SetenvCommand) Run(args []string) int {
	loginName := args[0]
	if args[1] != "exec" {
		log.Println("$ developers-account-mapper setenv <user> exec")
		return 1
	}

	s := store.NewDynamoDB()

	user, err := s.GetUserByLoginName(loginName)
	if err != nil {
		log.Println(err)
		return 1
	}
	os.Setenv("GITHUB_USERNAME", user.GitHubUsername)

	return 0
}

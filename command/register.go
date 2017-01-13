package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/developers-account-mapper/models"
	"github.com/wantedly/developers-account-mapper/store"
)

type RegisterCommand struct {
	Meta
}

func (c *RegisterCommand) Run(args []string) int {
	var loginName, githubUsername, slackUsername string
	if len(args) == 3 {
		loginName = args[0]
		githubUsername = args[1]
		slackUsername = args[2]
	} else {
		log.Println(c.Help())
		return 1
	}

	s := store.NewDynamoDB()

	user := models.NewUser(loginName, githubUsername, slackUsername, "")
	err := user.RetrieveSlackUserId()
	if err != nil {
		log.Println(err)
		return 1
	}
	if err := s.AddUser(user); err != nil {
		log.Println(err)
		return 1
	}

	userSummary, err := user.String()
	if err != nil {
		log.Println(err)
		return 1
	}
	fmt.Printf("user %q added.\n", userSummary)

	return 0
}

func (c *RegisterCommand) Synopsis() string {
	return "Register LoginName and other accounts mapping"
}

func (c *RegisterCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

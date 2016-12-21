package command

import (
	"strings"
	"os"
	"log"
	"fmt"

	"github.com/wantedly/developers-account-mapper/store"
	"github.com/wantedly/developers-account-mapper/models"
)

type RegisterCommand struct {
	Meta
}

func (c *RegisterCommand) Run(args []string) int {
	var loginName, githubUsername, slackUsername string
	if len(args) == 2 {
		loginName = os.Getenv("USER")
		githubUsername = args[0]
		slackUsername  = args[1]
	} else if len(args) == 3 {
		loginName = args[0]
		githubUsername = args[1]
		slackUsername = args[2]
	} else {
		log.Println(c.Help())
		return 1
	}

	s := store.NewDynamoDB()

	user := models.NewUser(loginName, githubUsername, slackUsername, "")
	user, err := user.RetrieveSlackUserId()
	if err != nil {
		log.Println(err)
		return 1
	}
	if err := s.AddUser(user); err != nil {
		log.Println(err)
		return 1
	}
	fmt.Printf("user %q added.\n", user)

	return 0
}

func (c *RegisterCommand) Synopsis() string {
	return "Register LoginName and GitHubUsername mapping"
}

func (c *RegisterCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

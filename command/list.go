package command

import (
	"strings"
	"github.com/wantedly/github-username-converter/store"
	"log"
	"fmt"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	s := store.NewDynamoDB()
	users, err := s.ListUsers()

	if err != nil {
		log.Println(err)
		return 1
	}
	for _, user := range users {
		fmt.Println(user)
	}

	return 0
}

func (c *ListCommand) Synopsis() string {
	return "List mapping of <login_name> and <github_username>"
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

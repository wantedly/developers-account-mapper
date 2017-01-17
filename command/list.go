package command

import (
	"log"
	"strings"

	"github.com/wantedly/developers-account-mapper/models"
	"github.com/wantedly/developers-account-mapper/store"
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

	models.PrintUsers(users)

	return 0
}

func (c *ListCommand) Synopsis() string {
	return "List mapping of <login_name> and mapped accounts"
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

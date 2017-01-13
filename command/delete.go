package command

import (
	"log"
	"strings"

	"github.com/wantedly/developers-account-mapper/store"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
	loginName := args[0]

	s := store.NewDynamoDB()
	err := s.DeleteUser(loginName)

	if err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

func (c *DeleteCommand) Synopsis() string {
	return "Delete record with <login_name>"
}

func (c *DeleteCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

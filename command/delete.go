package command

import (
	"strings"

	"github.com/wantedly/developers-account-mapper/store"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
	loginName := args[0]

	s := store.NewDynamoDB()
	s.DeleteUser(loginName)

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

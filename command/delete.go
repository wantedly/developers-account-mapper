package command

import (
	"strings"
)

type DeleteCommand struct {
	Meta
}

func (c *DeleteCommand) Run(args []string) int {
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

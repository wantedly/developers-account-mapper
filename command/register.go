package command

import (
	"strings"
)

type RegisterCommand struct {
	Meta
}

func (c *RegisterCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *RegisterCommand) Synopsis() string {
	return ""
}

func (c *RegisterCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

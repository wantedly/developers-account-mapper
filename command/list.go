package command

import (
	"strings"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	// Write your code here

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

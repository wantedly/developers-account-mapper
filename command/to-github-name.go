package command

import (
	"strings"
)

type ToGithubNameCommand struct {
	Meta
}

func (c *ToGithubNameCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ToGithubNameCommand) Synopsis() string {
	return ""
}

func (c *ToGithubNameCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

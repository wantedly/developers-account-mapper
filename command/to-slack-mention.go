package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/developers-account-mapper/store"
)

type ToSlackMention struct {
	Meta
}

func (c *ToSlackMention) Run(args []string) int {
	var loginName string
	if len(args) == 1 {
		loginName = args[0]
	} else {
		log.Println(c.Help())
		return 1
	}

	s := store.NewDynamoDB()

	user, err := s.GetUserByLoginName(loginName)
	if err != nil {
		log.Println(err)
		return 1
	}
	fmt.Printf("SlackMention account for %s is: %s\n", loginName, user.SlackMention())

	return 0
}

func (c *ToSlackMention) Synopsis() string {
	return "Get <slack_mention> from <login_name>"
}

func (c *ToSlackMention) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

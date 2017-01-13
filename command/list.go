package command

import (
	"fmt"
	"log"
	"strings"

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

	for _, user := range users {
		userSummary, err := user.String()
		if err != nil {
			log.Println(err)
			return 1
		}
		fmt.Println(userSummary)
	}

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

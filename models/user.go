package models

import (
	"fmt"

	"github.com/wantedly/developers-account-mapper/services"
)

// User stores login user name and accounts information
type User struct {
	LoginName      string
	GitHubUsername string
	SlackUsername  string
	SlackUserId    string
}

// NewUser creates new User instance
func NewUser(loginName string, githubUsername string, slackUsername string, slackUserId string) *User {
	return &User{
		LoginName:      loginName,
		GitHubUsername: githubUsername,
		SlackUsername:  slackUsername,
		SlackUserId:    slackUserId,
	}
}

const Headers = []string{
	"AWS IAM",
	"GITHUB",
	"SLACK",
}

func (u *User) Envs() []string {
	return []string{
		fmt.Sprintf("GITHUB_USERNAME=%s", u.GitHubUsername),
		fmt.Sprintf("SLACK_MENTION=%s", u.SlackMention()),
	}
}

func (u *User) SlackMention() string {
	return fmt.Sprintf("<@%v|%v>", u.SlackUserId, u.SlackUsername)
}

func (u *User) RetrieveSlackUserId() error {
	nameIdMap, err := services.SlackUserList()
	if err != nil {
		return err
	}
	u.SlackUserId = nameIdMap[u.SlackUsername]
	return nil
}

func (u *User) String() (string, error) {
	if u.SlackUserId == "" {
		err := u.RetrieveSlackUserId()
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%v:@%v:%v", u.LoginName, u.GitHubUsername, u.SlackMention()), nil
}

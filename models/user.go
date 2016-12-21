package models

import (
	"fmt"

	"github.com/wantedly/developers-account-mapper/services"
)

// User stores login user name and GitHub username
type User struct {
	LoginName      string
	GitHubUsername string
	SlackUsername  string
	SlackUserId    string
}

// NewUser creates new User instance
func NewUser(loginName string, githubUsername string, slackUsername string, slackUserId string) *User {
	return &User{
		LoginName: loginName,
		GitHubUsername: githubUsername,
		SlackUsername: slackUsername,
		SlackUserId: "",
	}
}

func (u *User) SetSlackUserId(newId string) *User {
	u.SlackUserId = newId
	return u
}

func (u *User) RetrieveSlackUserId() (*User, error) {
	nameIdMap, err := services.SlackUserList()
	if err != nil {
		return nil, err
	}
	u.SetSlackUserId(nameIdMap[u.SlackUsername])
	return u, nil
}

func (u *User) String() string {
	if u.SlackUserId == "" {
		u.RetrieveSlackUserId()
	}
	return fmt.Sprintf("%v:@%v:<@%v:%v>", u.LoginName, u.GitHubUsername, u.SlackUsername, u.SlackUserId)
}

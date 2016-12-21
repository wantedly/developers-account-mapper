package models

import (
	"fmt"
)

// User stores login user name and GitHub username
type User struct {
	LoginName string
	GitHubUsername string
	SlackUsername string
}

// NewUser creates new User instance
func NewUser(loginName string, githubUsername string, slackUsername string) *User {
	return &User{
		LoginName: loginName,
		GitHubUsername: githubUsername,
		SlackUsername: slackUsername,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("%v:@%v", u.LoginName, u.GitHubUsername)
}

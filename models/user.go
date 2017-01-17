package models

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

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

var userHeaders = []string{
	"AWS-IAM",
	"GITHUB",
	"SLACK",
	"SLACK_MENTION",
}

func PrintUsers(users []*User) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, strings.Join(userHeaders, "\t"))
	for _, user := range users {
		fmt.Fprintln(w, strings.Join(user.Attributes(), "\t"))
	}
	w.Flush()
}

func (u *User) Attributes() []string {
	return []string{
		u.LoginName,
		u.GitHubUsername,
		u.SlackUsername,
		u.SlackMention(),
	}
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

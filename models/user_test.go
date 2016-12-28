package models

import (
	"testing"
)

func SetUser() *User {
	return &User {
		LoginName: "loginName",
		GitHubUsername: "github_user",
		SlackUsername: "slack_user",
		SlackUserId: "SLACKID",
	}
}

func TestSlackMention(t *testing.T) {
	user := SetUser()

	expect := "<@SLACKID|slack_user>"
	actual := user.SlackMention()

	if actual != expect {
		t.Fatalf("%v does not much to expected: %v", actual, expect)
	}
}

package models

import (
	"testing"
)

func TestSlackMention(t *testing.T) {
	user := &User{
		LoginName:      "loginName",
		GitHubUsername: "github_user",
		SlackUsername:  "slack_user",
		SlackUserId:    "SLACKID",
	}

	expect := "<@SLACKID|slack_user>"
	actual := user.SlackMention()

	if actual != expect {
		t.Fatalf("%v does not much to expected: %v", actual, expect)
	}
}

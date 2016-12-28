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

func TestEnvs(t *testing.T) {
	user := SetUser()

	expect := []string{
		"GITHUB_USERNAME=github_user",
		"SLACK_MENTION=<@SLACKID|slack_user>",
	}
	actual := user.Envs()

	if len(actual) != len(expect) {
		t.Fatalf("%v does not much to expected: %v", len(actual), len(expect))
	}

	for i := range expect {
		if expect[i] != actual[i] {
			t.Fatalf("%v does not much to expected: %v", actual, expect)
		}
	}
}

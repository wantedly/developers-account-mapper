package models

import (
	"reflect"
	"testing"
)

func SetUser() *User {
	return &User{
		LoginName:      "loginName",
		GitHubUsername: "github_user",
		SlackUsername:  "slack_user",
		SlackUserId:    "SLACKID",
	}
}

func TestSlackMention(t *testing.T) {
	user := SetUser()

	expect := "<@SLACKID|slack_user>"
	actual := user.SlackMention()

	if actual != expect {
		t.Fatalf("%v does not match to expected: %v", actual, expect)
	}
}

func TestEnvs(t *testing.T) {
	user := SetUser()

	expect := []string{
		"GITHUB_USERNAME=github_user",
		"SLACK_MENTION=<@SLACKID|slack_user>",
	}
	actual := user.Envs()

	if !reflect.DeepEqual(actual, expect) {
		t.Fatalf("%v does not match to expected: %v", actual, expect)
	}
}

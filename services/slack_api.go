package services

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/jmoiron/jsonq"
)

const (
	slackUserListURL = "https://slack.com/api/users.list"
)

func SlackUserList() ([]interface{}){
	if token := os.Getenv("SLACK_API_TOKEN"); token == "" {
		return nil, fmt.Errorf("You need to pass SLACK_API_TOKEN as environment variable")
	}
	requestURL := slackUserListURL + "?token=" + token
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := map[string]interface{}{}
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	arr, err := jq.Array("members")
	return arr, err
}

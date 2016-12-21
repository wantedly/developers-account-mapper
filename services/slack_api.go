package services

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/jmoiron/jsonq"
	"strconv"
)

const (
	slackUserListURL = "https://slack.com/api/users.list"
)

func SlackUserList() (map[string]string, error){
	token := os.Getenv("SLACK_API_TOKEN")
	if token == "" {
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

	users := make(map[string]string)
	for i := 0; i < len(arr); i++ {
		id, _ := jq.String("members", strconv.Itoa(i), "id")
		name, _ := jq.String("members", strconv.Itoa(i), "name")
		users[name] = id
	}
	return users, err
}

package services

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/jmoiron/jsonq"
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
		id, idErr := jq.String("members", strconv.Itoa(i), "id")
		name, nameErr := jq.String("members", strconv.Itoa(i), "name")
		if idErr != nil && nameErr != nil {
			return nil, idErr || nameErr
		}
		users[name] = id
	}
	return users, err
}

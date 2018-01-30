package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jmoiron/jsonq"
)

const (
	slackUserListURL = "https://slack.com/api/users.list"
)

func SlackUserList(token string) (map[string]string, error) {
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
		id, err := jq.String("members", strconv.Itoa(i), "id")
		if err != nil {
			return nil, err
		}
		name, err := jq.String("members", strconv.Itoa(i), "name")
		if err != nil {
			return nil, err
		}
		users[name] = id
	}
	return users, err
}

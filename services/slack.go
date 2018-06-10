package services

import (
	"bytes"
	"drink-api/config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// SlackPayload is slack payload for sending message.
type SlackPayload struct {
	Text     string `json:"text,omitempty"`
	UserName string `json:"username,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

const (
	DefaultUserName = "yuile"
)

// SendMessage send a message to channel
func SendMessage(msg string, channel string) (err error) {
	config.Load()
	slackURL := config.MustGet("SLACK_WEBHOOK_URL")

	payload, _ := json.Marshal(&SlackPayload{
		Text:     msg,
		UserName: DefaultUserName,
	})

	fmt.Println(string(payload[:]))

	request, err := http.NewRequest(
		"POST",
		slackURL,
		bytes.NewReader(payload),
	)

	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if string(body[:]) != "ok" {
		msg := fmt.Sprintf("can not send message. error: %s\n", string(body[:]))
		return errors.New(msg)
	}

	if err != nil {
		return err
	}

	return nil
}

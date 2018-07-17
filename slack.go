package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	AsUser  bool   `json:"as_user"`
}

type Slack struct {
	token string
}

func New() *Slack {
	return &Slack{
		token: os.Getenv("SLACK_TOKEN"),
	}
}

func (s *Slack) TypeID() string {
	return "slack"
}

func (s *Slack) Send(token map[string]string, message string) error {

	ms := &Message{
		Channel: token["slackChannelID"],
		Text:    message,
		AsUser:  true,
	}

	m, err := json.Marshal(ms)
	if err != nil {
		return err
	}

	body := bytes.NewReader(m)
	req, err := http.NewRequest("POST", "https://slack.com/api/chat.postMessage", body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

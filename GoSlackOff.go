package goslackoff

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type _GoSlackOff struct {
	Domain   string
	Token    string
	Username string
	Channel  string
}

type Message struct {
	Text     string `json:"text"`
	Channel  string `json:"channel"`
	Username string `json:"username"`
}

func New(domain string, token string, username string, channel string) _GoSlackOff {
	return _GoSlackOff{
		Domain:   domain,
		Token:    token,
		Username: username,
		Channel:  channel,
	}
}

func (this _GoSlackOff) SendMessage(message Message) (bool, *http.Response) {
	fmt.Println("Sending message!")

	//-----------------------------------------
	// Inherit stuff
	//-----------------------------------------
	if message.Channel == "" {
		message.Channel = this.Channel
	}
	if message.Username == "" {
		message.Username = this.Username
	}

	url := fmt.Sprintf("https://%s.slack.com/services/hooks/incoming-webhook?token=%s", this.Domain, this.Token)
	jsonStr, _ := json.Marshal(message)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) == "ok" {
		return true, resp
	}

	return false, resp
}

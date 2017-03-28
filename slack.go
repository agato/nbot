package nbot

import (
	"net/http"
	"net/url"
	"fmt"
	"bytes"
	"io/ioutil"
)

const (
	SLACK_BASE_URL = "https://slack.com"
	SLACK_CHAT_URL = "/api/chat.postMessage"
)

type SlackApi struct {
	Token string
	User string
	HttpClient           *http.Client
}

func NewSlackApi(token string, user string) *SlackApi {
	return &SlackApi{Token: token, User: user, HttpClient: http.DefaultClient}
}

func (c *SlackApi) SendMessage(channelid string, msg string) error {
	u, _ := url.ParseRequestURI(SLACK_BASE_URL)
	u.Path = SLACK_CHAT_URL
	reqUrl := fmt.Sprintf("%v", u)

	body := url.Values{}
	body.Set("token", c.Token)
	if len(c.User) > 0 {
		body.Set("username", c.User)
	}
	body.Set("channel", channelid)
	body.Set("text", msg)

	r, _ := http.NewRequest("POST", reqUrl, bytes.NewBufferString(body.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := c.HttpClient.Do(r)

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Printf("result: %s\n", contents)
	return nil
}
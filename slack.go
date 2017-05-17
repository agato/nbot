package nbot

import (
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"bytes"
	"io/ioutil"
)

const (
	SLACK_BASE_URL = "https://slack.com"
	SLACK_CHAT_URL = "/api/chat.postMessage"
	SLACK_INCOM_HOOK_URL = "https://hooks.slack.com/services/%s/%s/%s"
)

type SlackApi struct {
	Token string
	User string
	HttpClient           *http.Client
}

type SlackIncomHookApi struct {
	TParam string
	BParam string
	Token string
	HttpClient           *http.Client
}

type SlackIncomHookMsgJson struct {
	Channel   string `json:"channel"`
	Text      string `json:"text"`
}

func NewSlackApi(token string, user string) *SlackApi {
	return &SlackApi{Token: token, User: user, HttpClient: http.DefaultClient}
}

func NewSlackIncomHookApi(tParam string, bParam string, token string) *SlackIncomHookApi {
	return &SlackIncomHookApi{TParam: tParam, BParam: bParam, Token: token, HttpClient: http.DefaultClient}
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

func (c *SlackIncomHookApi) SendMessage(channelid string, msg string) error {
	u, _ := url.ParseRequestURI(fmt.Sprintf(SLACK_INCOM_HOOK_URL, c.TParam, c.BParam, c.Token))
	reqUrl := fmt.Sprintf("%v", u)

	body := url.Values{}
	msgJson, _ := json.Marshal(SlackIncomHookMsgJson{channelid, fmt.Sprintf("%s", msg)})
	body.Set("payload", string(msgJson))

	r, _ := http.NewRequest("POST", reqUrl, bytes.NewBufferString(body.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := c.HttpClient.Do(r)

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Printf("result: %s\n", contents)
	return nil
}

package nbot

import (
	"net/http"
	"net/url"
	"fmt"
	"bytes"
	"io/ioutil"
)

const (
	CHATWORK_BASE_URL = "https://api.chatwork.com"
	CHATWORK_ROOM_URL = "/v2/rooms/%s/messages"
)

type ChatWorkApi struct {
	Token string
	HttpClient           *http.Client
}

func NewChatWorkApi(token string) *ChatWorkApi {
	return &ChatWorkApi{Token: token, HttpClient: http.DefaultClient}
}

func (c *ChatWorkApi) SendMessage(roomid string, msg string) error {
	u, _ := url.ParseRequestURI(CHATWORK_BASE_URL)
	u.Path = fmt.Sprintf(CHATWORK_ROOM_URL, roomid)
	reqUrl := fmt.Sprintf("%v", u)

	body := url.Values{}
	body.Set("body", msg)

	r, _ := http.NewRequest("POST", reqUrl, bytes.NewBufferString(body.Encode()))
	r.Header.Add("X-ChatWorkToken", c.Token)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := c.HttpClient.Do(r)

	defer resp.Body.Close()
	contents, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Printf("result: %s\n", contents)
	return nil
}
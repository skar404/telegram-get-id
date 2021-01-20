package telegram

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/skar404/telegram-get-id/telegram/object"
)

type Config struct {
	// default: "https://api.telegram.org/bot%s/"
	Url string
}

func (c *Config) getUrl(s string) string {
	return c.Url + s
}

func (c *Config) httpClient(method string, url string, jsonBody map[string]interface{}, object interface{}) error {
	byteBody, err := json.Marshal(jsonBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(byteBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	BodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(BodyRaw, &object)

	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetUpdates(offSet int) (object.GetUpdate, error) {
	url := c.getUrl("getUpdates")

	jsonBody := make(map[string]interface{})
	if offSet != 0 {
		jsonBody["offset"] = strconv.Itoa(offSet)
	}

	resUpdate := object.GetUpdate{}
	err := c.httpClient("POST", url, jsonBody, &resUpdate)

	return resUpdate, err
}

func (c *Config) DeleteWebhook() error {
	url := c.getUrl("deleteWebhook")
	err := c.httpClient("POST", url, nil, nil)
	return err
}

func (c *Config) SendMessage(chatId int, text string) error {
	url := c.getUrl("sendMessage")
	jsonBody := map[string]interface{}{
		"chat_id": chatId,
		"text":    text,
	}

	return c.httpClient("POST", url, jsonBody, nil)
}

func (c *Config) SetChatDescription(chatId int, text string) error {
	url := c.getUrl("setChatDescription")
	jsonBody := map[string]interface{}{
		"chat_id":     chatId,
		"description": text,
	}

	return c.httpClient("POST", url, jsonBody, nil)
}

func (c *Config) GetMe() (object.GetMe, error) {
	url := c.getUrl("getMe")

	resUpdate := object.GetMe{}

	err := c.httpClient("GET", url, nil, &resUpdate)
	return resUpdate, err
}

func (c *Config) SetWebHook(hookUrl string, maxConn int) error {
	url := c.getUrl("setWebhook")
	jsonBody := map[string]interface{}{
		"url": hookUrl,
	}

	if maxConn != 0 {
		jsonBody["max_connections"] = maxConn
	}

	return c.httpClient("POST", url, jsonBody, nil)
}

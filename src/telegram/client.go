package telegram

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"telegram-get-id/src/telegram/object"
)

type Config struct {
	// default: "https://api.telegram.org/bot%s/"
	Url string
}

func (c *Config) getUrl(s string) string {
	return c.Url + s
}

func (c *Config) httpClient(method string, url string, jsonBody map[string]interface{}, object interface{}) {
	byteBody, err := json.Marshal(jsonBody)
	if err != nil {
		// TODO add logs
		log.Fatal("")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(byteBody))
	if err != nil {
		// TODO add logs
		log.Fatal("NewRequest")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		// TODO add logs
		log.Fatal("Do")
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&object)

	if err != nil {
		// TODO add logs
		log.Fatal("NewDecoder")
	}
}

func (c *Config) GetUpdates(offSet int) object.GetUpdate {
	url := c.getUrl("getUpdates")

	jsonBody := make(map[string]interface{})
	if offSet != 0 {
		jsonBody["offset"] = strconv.Itoa(offSet)
	}

	resUpdate := object.GetUpdate{}
	c.httpClient("POST", url, jsonBody, &resUpdate)

	return resUpdate
}

func (c *Config) SendMessage(chatId int, text string) {
	url := c.getUrl("sendMessage")
	jsonBody := map[string]interface{}{
		"chat_id":    chatId,
		"text":       text,
		"parse_mode": "Markdown",
	}

	c.httpClient("POST", url, jsonBody, nil)
}

func (c *Config) SetChatDescription(chatId int, text string) {
	url := c.getUrl("setChatDescription")
	jsonBody := map[string]interface{}{
		"chat_id":     chatId,
		"description": text,
	}

	c.httpClient("POST", url, jsonBody, nil)
}

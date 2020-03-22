package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *Config) GetUpdates(offSet int) {
	url := c.getUrl("getUpdates")

	resBody := make(map[string]string)
	if offSet != 0 {
		resBody["offSet"] = strconv.Itoa(offSet)
	}

	jsonBody, _ := json.Marshal(resBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		// TODO add logs
		log.Fatal("")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		// TODO add logs
		log.Fatal("")
	}

	defer resp.Body.Close()

	respInt := object.GetUpdate{}
	_ = json.NewDecoder(resp.Body).Decode(&respInt)

	fmt.Println(respInt)
}

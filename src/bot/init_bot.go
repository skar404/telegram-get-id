package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"telegram-get-id/src/telegram"
	"telegram-get-id/src/telegram/object"
	"telegram-get-id/src/utils"
)

type Config struct {
	BotToken    string
	Mod         string
	TelegramUrl string

	AppHost string
	AppPort string

	Debug bool

	botKey  string
	botPath string

	tgClient telegram.Config
}

func (c *Config) init() error {
	if c.BotToken == "" {
		return errors.New("token is nil")
	}

	if c.TelegramUrl == "" {
		c.TelegramUrl = fmt.Sprintf("https://api.telegram.org/bot%s/", c.BotToken)
	} else {
		c.TelegramUrl = fmt.Sprintf(c.TelegramUrl, c.BotToken)
	}

	// init Telegram Client
	c.tgClient = telegram.Config{
		Url: c.TelegramUrl,
	}

	c.botKey = utils.RandStringRunes(50)
	c.botPath = "telegram/" + c.botKey

	return c.isValidToken()
}

func (c *Config) isValidToken() error {
	bot, err := c.tgClient.GetMe()

	if err != nil {
		log.Println(err)
		return errors.New("connect fail")
	} else if bot.Ok == false {
		return errors.New("not valid token")
	}
	return nil
}

func getMessage(item object.Update) (object.Message, error) {
	_empty := object.Message{}
	var message object.Message

	if item.Message != _empty {
		message = item.Message
	} else if item.EditedMessage != _empty {
		message = item.EditedMessage
	} else if item.ChannelPost != _empty {
		message = item.ChannelPost
	} else if item.EditedChannelPost != _empty {
		message = item.EditedChannelPost
	} else {
		return message, errors.New("not found messages")
	}
	return message, nil
}

func skipText(text string) bool {
	validMessage := []string{
		"id", "/id", "@id", "get id", "get_id", "get-id", "/get_id", "/get-id", "@send_id_ru_bot", "@send_id_bot",
		"@send_id",
	}

	return !utils.StringInSlice(strings.ToLower(text), validMessage)
}

func (c Config) sendIds(item object.Update) {
	message, err := getMessage(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	if skipText(message.Text) {
		return
	}

	chatId := message.Chat.Id

	log.Println(fmt.Sprintf("Send message to chat_id=%d update_id=%d", chatId, item.UpdateId))

	sendMessage := fmt.Sprintf("Chat ID: %d", chatId)

	err = c.tgClient.SendMessage(chatId, sendMessage)
	if err != nil {
		log.Println("error SendMessage message=" + sendMessage)
	}
	err = c.tgClient.SetChatDescription(chatId, sendMessage)
	if err != nil {
		log.Println("error SetChatDescription message=" + sendMessage)
	}
}

func (c *Config) GetUpdates() {
	updateId := 0

	for true {
		raw, err := c.tgClient.GetUpdates(updateId)

		if err != nil {
			log.Println(err)
			continue
		}

		for _, item := range raw.Result {
			c.sendIds(item)

			updateId = item.UpdateId + 1
		}
	}
}

func (c *Config) telegramWebHook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}

	update := object.Update{}

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Println("error decode body json")
		w.WriteHeader(400)
		return
	}

	c.sendIds(update)
	w.WriteHeader(200)
}

func (c *Config) WebHook() {
	// set
	hookLink := c.AppHost + c.botPath
	err := c.tgClient.SetWebHook(hookLink, 0)
	if err != nil {
		log.Fatalln(fmt.Sprintf("error set webhook=%s", hookLink))
	}
	log.Println(fmt.Sprintf("set webhook=%s", hookLink))

	// init web server
	http.HandleFunc("/"+c.botPath, c.telegramWebHook)

	// set host
	host := ""
	if c.Debug == true {
		host = "127.0.0.1"
	}

	host += ":" + c.AppPort
	if c.AppPort == "" {
		host += ":8080"
	}

	log.Println("Start web app host: " + host)
	log.Fatal(http.ListenAndServe(host, nil))
}

func (c *Config) Start() error {
	err := c.init()
	if err != nil {
		return err
	}

	if c.Mod == "GET_UPDATES" {
		c.GetUpdates()
	} else if c.Mod == "WEB_HOOK" {
		c.WebHook()
	}
	log.Fatalln("not valid mod")

	return nil
}

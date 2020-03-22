package bot

import (
	"fmt"
	"log"

	"telegram-get-id/src/telegram"
	"telegram-get-id/src/telegram/object"
)

type Config struct {
	BotToken    string
	Mod         string
	TelegramUrl string

	tgClient telegram.Config
}

func (c *Config) init() {
	if c.TelegramUrl == "" {
		c.TelegramUrl = fmt.Sprintf("https://api.telegram.org/bot%s/", c.BotToken)
	} else {
		c.TelegramUrl = fmt.Sprintf(c.TelegramUrl, c.BotToken)
	}

	// init Telegram Client
	c.tgClient = telegram.Config{
		Url: c.TelegramUrl,
	}
}

func getMessage(item object.Update) object.Message {
	_empty := (object.Message{})
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
		// TODO add logs
		log.Fatal("ERRR")
	}
	return message
}

func (c *Config) sendIds() {
	updateId := 0
	for true {
		fmt.Println(fmt.Sprintf("updateId=%d", updateId))
		raw := c.tgClient.GetUpdates(updateId)

		for _, item := range raw.Result {
			message := getMessage(item)

			fmt.Println(fmt.Sprintf("Send message to chat_id=%d update_id=%d", message.Chat.Id, item.UpdateId))
			c.tgClient.SendMessage(message.Chat.Id, fmt.Sprintf("Chat ID: %d", message.Chat.Id))

			updateId = item.UpdateId + 1
		}
	}
}

func (c *Config) Start() {
	c.init()
	if c.Mod == "GET_UPDATES" {
		c.sendIds()
	}
}

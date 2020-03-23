package bot

import (
	"errors"
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

func (c Config) sendIds(item object.Update) {
	message, err := getMessage(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	chatId := message.Chat.Id

	fmt.Println(fmt.Sprintf("Send message to chat_id=%d update_id=%d", chatId, item.UpdateId))

	sendMessage := fmt.Sprintf("Chat ID: %d", chatId)

	err = c.tgClient.SendMessage(chatId, sendMessage)
	if err != nil {
		fmt.Println("error SendMessage message=" + sendMessage)
	}
	err = c.tgClient.SetChatDescription(chatId, sendMessage)
	if err != nil {
		fmt.Println("error SetChatDescription message=" + sendMessage)
	}
}

func (c *Config) GetUpdates() {
	updateId := 0

	for true {
		raw, err := c.tgClient.GetUpdates(updateId)

		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, item := range raw.Result {
			c.sendIds(item)

			updateId = item.UpdateId + 1
		}
	}
}

func (c *Config) WebHook() {

}

func (c *Config) Start() {
	c.init()
	if c.Mod == "GET_UPDATES" {
		c.GetUpdates()
	} else if c.Mod == "WEB_HOOK" {
		c.WebHook()
	}
	log.Fatalln("not valid mod")
}

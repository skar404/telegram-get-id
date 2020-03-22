package bot

import (
	"fmt"

	"telegram-get-id/src/telegram"
)

type Config struct {
	BotToken    string
	Mod         string
	TelegramUrl string
}

func (c *Config) init() {
	if c.TelegramUrl == "" {
		c.TelegramUrl = fmt.Sprintf("https://api.telegram.org/bot%s/", c.BotToken)
	} else {
		c.TelegramUrl = fmt.Sprintf(c.TelegramUrl, c.BotToken)
	}
}

func (c *Config) Start() {
	c.init()

	// init Telegram Client
	tg := telegram.Config{
		Url: c.TelegramUrl,
	}

	if c.Mod == "GET_UPDATES" {
		tg.GetUpdates(0)
	}
}

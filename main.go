package main

import (
	"fmt"
	"log"
	"os"

	"telegram-get-id/src/bot"
	"telegram-get-id/src/utils"
)

type Setting struct {
	Mod         string
	BotToken    string
	TelegramUrl string
}

func initEnv() Setting {
	// Init
	s := Setting{
		Mod:      os.Getenv("MOD"),
		BotToken: os.Getenv("BOT_TOKEN"),
	}

	// Validate:
	if !utils.StringInSlice(s.Mod, []string{"WEB_HOOK", "GET_UPDATES"}) {
		log.Fatalln(fmt.Sprintf("Not Valid env MOD=%s", s.Mod))
	}

	return s
}

func main() {
	env := initEnv()
	log.Println(fmt.Sprintf("Start app, mod: %s", env.Mod))

	myBot := bot.Config{
		Mod:      env.Mod,
		BotToken: env.BotToken,
	}

	err := myBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}

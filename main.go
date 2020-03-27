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

	AppHost string

	Debug bool
}

func initEnv() Setting {
	// Init
	s := Setting{
		Mod:      os.Getenv("MOD"),
		BotToken: os.Getenv("BOT_TOKEN"),

		AppHost: os.Getenv("APP_HOST"),
	}

	if utils.StringInSlice(os.Getenv("DEBUG"), []string{"True", "true", "1"}) {
		s.Debug = true
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
		AppHost:  env.AppHost,

		Debug: env.Debug,
	}

	err := myBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}

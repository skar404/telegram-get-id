package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/skar404/telegram-get-id/bot"
	"github.com/skar404/telegram-get-id/utils"
)

type Setting struct {
	Mod   string
	Debug bool
}

func initEnv() (Setting, error) {
	// Init
	s := Setting{
		Mod: os.Getenv("MOD"),
	}

	if utils.StringInSlice(os.Getenv("DEBUG"), []string{"True", "true", "1"}) {
		s.Debug = true
	}

	// Validate:
	if !utils.StringInSlice(s.Mod, []string{"WEB_HOOK", "GET_UPDATES"}) {
		return s, errors.New(fmt.Sprintf("not Valid env MOD=%s", s.Mod))
	}
	return s, nil
}

func main() {
	env, err := initEnv()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(fmt.Sprintf("Start app, mod: %s", env.Mod))

	myBot := bot.Config{
		Mod:      env.Mod,
		BotToken: os.Getenv("BOT_TOKEN"),
		AppHost:  os.Getenv("APP_HOST"),
		AppPort:  os.Getenv("PORT"),

		Debug: env.Debug,
	}

	err = myBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}

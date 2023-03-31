package main

import (
	"github.com/alekseyprokopov/bot256/internal/app/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")

	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//
	bot.Debug = true

	log.Printf("Authorized on account: %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	r := router.New(bot)

	for update := range updates {
		r.HandleUpdate(update)
	}

}

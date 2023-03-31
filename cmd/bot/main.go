package main

import (
	"github.com/alekseyprokopov/bot256/internal/app/commander"
	"github.com/alekseyprokopov/bot256/internal/app/router"
	"github.com/alekseyprokopov/bot256/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	t := os.Getenv("TOKEN")

	productService := product.NewService()
	bot, err := tgbotapi.NewBotAPI(t)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account: %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	commander := commander.New(bot, productService)
	r := router.New(commander, productService)

	for update := range updates {
		r.HandleUpdate(update)
	}

}

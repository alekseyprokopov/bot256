package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	t:=os.Getenv("TOKEN")

	bot, err:=tgbotapi.NewBotAPI(t)
	if err!=nil{
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account: %s", bot.Self.UserName)


}

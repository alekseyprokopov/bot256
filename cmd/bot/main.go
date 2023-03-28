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

	u:=tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err:= bot.GetUpdatesChan(u)
	if err!=nil{
		log.Panic(err)
	}

	for update := range updates{
		if update.Message == nil{
			continue
		}
		log.Printf("[FROM: %s],TEXT: %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "echo: " + update.Message.Text)
		bot.Send(msg)
	}

}

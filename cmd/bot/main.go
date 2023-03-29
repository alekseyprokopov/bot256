package main

import (
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

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCmd(bot, update.Message)
		case "list":
			listCmd(bot, update.Message, productService)
		default:
			defCmd(bot, update.Message)
		}

		log.Printf("[FROM: %s],TEXT: %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "echo: "+update.Message.Text)
		bot.Send(msg)
	}

}

func helpCmd(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	text := "/help - help\n" +
		"/list - list products"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	bot.Send(msg)
}

func listCmd(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service){
	text:= "all products: \n"
	products := productService.List()

	for _, item := range products{
		text+= item.Name
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	bot.Send(msg)

}


func defCmd(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	text := "/help - help\n" +
		"/list - list products"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	bot.Send(msg)
}
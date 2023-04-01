package product

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) Default(message *tgbotapi.Message) {
	text := "default behaviour\n"
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	c.Bot.Send(msg)
}

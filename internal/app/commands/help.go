package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) HelpCmd(inputMessage *tgbotapi.Message) {
	text := "/help - help\n" +
		"/list - list products"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	c.Bot.Send(msg)
}

package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) DefCmd(inputMessage *tgbotapi.Message) {
	text := "default behaviour\n"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	c.Bot.Send(msg)
}

package product

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) Help(message *tgbotapi.Message) {
	text := "/help - help\n" +
		"/get - get product\n" +
		"/list - list products\n" +
		"/delete - delete product\n" +
		"/new - new product\n" +
		"/edit - update product\n"
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	c.Bot.Send(msg)
}

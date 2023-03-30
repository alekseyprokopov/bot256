package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) ListCmd(inputMessage *tgbotapi.Message) {
	text := "all products: \n"
	products := c.ProductService.List()

	for _, item := range products {
		text += item.Name + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	c.Bot.Send(msg)

}

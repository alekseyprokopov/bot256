package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(message *tgbotapi.Message) {
	text := "list of all products: \n\n"
	products, _ := c.ProductService.List(0, 10)

	for _, item := range *products {
		text += item.Name + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	c.Bot.Send(msg)

}

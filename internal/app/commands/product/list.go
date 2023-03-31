package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) ListCmd(inputMessage *tgbotapi.Message) {
	text := "all products: \n"
	products, _ := c.ProductService.List(0, 10)

	for _, item := range *products {
		text += item.Name + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	c.Bot.Send(msg)

}

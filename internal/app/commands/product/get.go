package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()
	id, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "can't get index: "+args)
		c.Bot.Send(msg)
		return
	}

	product, err := c.ProductService.Get(int64(id))
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "can't get product: "+string(id))
		c.Bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Product: %+v", product))
	c.Bot.Send(msg)
}

package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *Commander) getCmd(callback *tgbotapi.CallbackQuery, data string) {
	id, err := strconv.Atoi(data)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "can't get index: "+data)
		c.Bot.Send(msg)
		return
	}

	product, err := c.ProductService.Get(int64(id))
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "can't get product: "+string(id))
		c.Bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, fmt.Sprintf("Product: %+v", product))
	c.Bot.Send(msg)
}

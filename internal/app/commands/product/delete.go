package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *Commander) deleteCmd(callback *tgbotapi.CallbackQuery, data string) {
	id, err := strconv.Atoi(data)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "can't get index: "+data)
		c.Bot.Send(msg)
		return
	}

	ok, err := c.ProductService.Remove(int64(id))
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "can't remove product: "+string(id))
		c.Bot.Send(msg)
		return
	}
	if ok {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, fmt.Sprintf("SUCCESS!"))
		c.Bot.Send(msg)
	}

}

package product

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *Commander) deleteCmd(message *tgbotapi.Message) {
	args := message.CommandArguments()
	id, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "can't get index: "+args)
		c.Bot.Send(msg)
		return
	}

	ok, err := c.ProductService.Remove(int64(id))
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "can't remove product: "+string(id))
		c.Bot.Send(msg)
		return
	}
	if ok {
		msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("SUCCESS!"))
		c.Bot.Send(msg)
	}

}

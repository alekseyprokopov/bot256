package product

import (
	"fmt"
	"github.com/alekseyprokopov/bot256/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *Commander) New(message *tgbotapi.Message) {
	args := message.CommandArguments()
	id, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "can't get index: "+args)
		c.Bot.Send(msg)
		return
	}

	item := product.Product{Name: args}

	idx, err := c.ProductService.Create(&item)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "can't remove product: "+string(id))
		c.Bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("SUCCESS! item id: %d", idx))
	c.Bot.Send(msg)

}

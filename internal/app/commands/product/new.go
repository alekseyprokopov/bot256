package product

import (
	"fmt"
	"github.com/alekseyprokopov/bot256/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *Commander) newCmd(callback *tgbotapi.CallbackQuery, data string) {
	id, err := strconv.Atoi(data)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "can't get index: "+data)
		c.Bot.Send(msg)
		return
	}

	item := product.Product{Name: data}

	idx, err := c.ProductService.Create(&item)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "can't remove product: "+string(id))
		c.Bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, fmt.Sprintf("SUCCESS! item id: %d", idx))
	c.Bot.Send(msg)

}

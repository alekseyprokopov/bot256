package commander

import (
	"github.com/alekseyprokopov/bot256/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	Bot            *tgbotapi.BotAPI
	ProductService *product.Service
}



func New(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		Bot:            bot,
		ProductService: service,
	}
}

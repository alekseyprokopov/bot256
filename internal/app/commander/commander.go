package commander

import (
	"github.com/alekseyprokopov/bot256/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"runtime/debug"
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

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		c.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		c.handleCommand(update.Message)
	}

}

func (c *Commander) handleCallback(callback *tgbotapi.CallbackQuery) {

}
func (c *Commander) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "help":
		c.HelpCmd(message)
	case "list":
		c.ListCmd(message)
	default:
		c.DefCmd(message)
	}

	log.Printf("[FROM: %s],TEXT: %s", message.From.UserName, message.Text)
}

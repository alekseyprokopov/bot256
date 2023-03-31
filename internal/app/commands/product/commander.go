package product

import (
	"github.com/alekseyprokopov/bot256/internal/app/path"
	"github.com/alekseyprokopov/bot256/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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

func (c *Commander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {

	//commandPath.CommandName
	switch commandPath.CommandName {
	case "help":
		c.HelpCmd(message)
	case "list":
		c.ListCmd(message)
	default:
		c.DefCmd(message)
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "get":
		c.getCmd(callback, callbackPath.CallbackData)
	case "new":
		c.newCmd(callback, callbackPath.CallbackData)
	case "delete":
		c.deleteCmd(callback, callbackPath.CallbackData)
	case "edit":
		c.editCmd(callback, callbackPath.CallbackData)
	default:

	}

}
func (c *Commander) handleCommand(message *tgbotapi.Message) {

	log.Printf("[FROM: %s],TEXT: %s", message.From.UserName, message.Text)
}

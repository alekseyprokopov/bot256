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
	case "get":
		c.getCmd(message)
	case "new":
		c.newCmd(message)
	case "delete":
		c.deleteCmd(message)
	case "edit":
		c.editCmd(message)
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	log.Printf("name: ", callbackPath.CallbackName)
	switch callbackPath.CallbackName {

	default:

	}

}
func (c *Commander) handleCommand(message *tgbotapi.Message) {

	log.Printf("[FROM: %s],TEXT: %s", message.From.UserName, message.Text)
}

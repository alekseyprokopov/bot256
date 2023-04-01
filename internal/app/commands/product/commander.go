package product

import (
	"github.com/alekseyprokopov/bot256/internal/app/path"
	"github.com/alekseyprokopov/bot256/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type ProductCommander interface {
	Help(message *tgbotapi.Message)
	Get(message *tgbotapi.Message)
	List(message *tgbotapi.Message)
	Delete(message *tgbotapi.Message)

	New(message *tgbotapi.Message)
	Edit(message *tgbotapi.Message)

	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
}

type Commander struct {
	Bot            *tgbotapi.BotAPI
	ProductService *product.Service
}

func NewProductCommander(bot *tgbotapi.BotAPI, service *product.Service) ProductCommander {
	return &Commander{
		Bot:            bot,
		ProductService: service,
	}
}

func (c *Commander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {

	//commandPath.CommandName
	switch commandPath.CommandName {
	case "help":
		c.Help(message)
	case "list":
		c.List(message)
	default:
		c.Default(message)
	case "get":
		c.Get(message)
	case "new":
		c.New(message)
	case "delete":
		c.Delete(message)
	case "edit":
		c.Edit(message)
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

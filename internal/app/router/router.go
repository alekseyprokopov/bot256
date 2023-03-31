package router

import (
	"github.com/alekseyprokopov/bot256/internal/app/commands/product"
	"github.com/alekseyprokopov/bot256/internal/app/path"
	product2 "github.com/alekseyprokopov/bot256/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"runtime/debug"
)

type Storage interface {
	Add()
	Get()
	Update()
	Delete()
}

type Router struct {
	bot     *tgbotapi.BotAPI
	product Commander
}
type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

func New(bot *tgbotapi.BotAPI) *Router {
	return &Router{
		bot:     bot,
		product: product.New(bot, product2.NewService()),
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		r.handleCallback(update.CallbackQuery)
	case update.Message != nil:
		r.handleMessage(update.Message)
	}

}

func (r *Router) handleCallback(callback *tgbotapi.CallbackQuery) {
	callbackPath, err := path.ParseCallback(callback.Data)
	if err != nil {
		log.Printf("Router.handleCallback: error parsing callback data `%s` - %v", callback.Data, err)
		return
	}

	switch callbackPath.Domain {
	case "product":
		r.product.HandleCallback(callback, callbackPath)
	}

}

func (r *Router) handleMessage(message *tgbotapi.Message) {
	if !message.IsCommand() {
		log.Printf("Router.handleMessage: message is not a command `%s`", message.Command())
		return
	}
	messagePath, err := path.ParseCommand(message.Command())
	if err != nil {
		log.Printf("Router.handleMessage: error parsing message data `%s` - %v", message.Text, err)
		return
	}

	switch messagePath.Domain {
	case "product":
		r.product.HandleCommand(message, messagePath)
	}

}

package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"runtime/debug"
)

type Commander interface {
	HandleCallback (callback *tgbotapi.CallbackQuery )
	HandleCommand (callback *tgbotapi.Message )
}

type Router struct {
	commander Commander
	bot *tgbotapi.BotAPI
}

func (r *Router) HandleUpdate(update tgbotapi.Update){
	defer func() {
		if panicValue:=recover(); panicValue!=nil{
			log.Printf("recovered from panic: %v\n%v", panicValue, string(debug.Stack()))
		}
	}()

	switch {
	case update.CallbackQuery != nil:
		r.commander.HandleCallback(update.CallbackQuery)
	case update.Message!=nil:
		r.commander.HandleCommand(update.Message)
	}

}
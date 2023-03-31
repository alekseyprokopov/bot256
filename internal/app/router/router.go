package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Storage interface {
	Add()
	Get()
	Update()
	Delete()
}

type Router struct {
	botCommander *Commander
	storage      *Storage
}
type Commander interface {
	HandleUpdate(update tgbotapi.Update)
}

func New(commander *Commander, sto *Storage) *Router {
	return &Router{
		botCommander: commander,
		storage:      sto,
	}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	(*r.botCommander).HandleUpdate(update)
}

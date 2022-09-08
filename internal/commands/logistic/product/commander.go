package commands

import (
	service "github/nuxxxcake/go-bot/internal/service/logistic/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *DummyProductCommander, inputMsg *tgbotapi.Message){}

type ProductCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type DummyProductCommander struct {
	bot     *tgbotapi.BotAPI
	service service.ProductService
}

func NewDummyProductCommander(bot *tgbotapi.BotAPI, service service.ProductService) *DummyProductCommander {
	return &DummyProductCommander{
		bot:     bot,
		service: service,
	}
}

func (c *DummyProductCommander) HandleUpdate(inputMsg *tgbotapi.Message) {
	switch inputMsg.Command() {
	case "help":
		c.Help(inputMsg)
	case "get":
		c.Get(inputMsg)
	case "list":
		c.List(inputMsg)
	case "new":
		c.New(inputMsg)
	case "delete":
		c.Delete(inputMsg)
	default:
		return
	}
}

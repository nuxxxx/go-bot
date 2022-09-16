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

	CbList(inputCb *tgbotapi.CallbackQuery)
}

type CallbackQueryItem struct {
	name           string
	additionalData interface{}
}

type DummyProductCommander struct {
	bot                *tgbotapi.BotAPI
	service            service.ProductService
	callbackQueryQueue []CallbackQueryItem
}

func NewCallbackQueryItem(name string, additionalData interface{}) *CallbackQueryItem {
	return &CallbackQueryItem{
		name:           name,
		additionalData: additionalData,
	}
}

func NewDummyProductCommander(bot *tgbotapi.BotAPI, service service.ProductService) *DummyProductCommander {
	return &DummyProductCommander{
		bot:                bot,
		service:            service,
		callbackQueryQueue: []CallbackQueryItem{},
	}
}

func (c *DummyProductCommander) HandleCallback(inputCb *tgbotapi.CallbackQuery) {
	if len(c.callbackQueryQueue) == 0 {
		return
	}

	lastCb := c.callbackQueryQueue[0]

	switch lastCb.name {
	case "list":
		c.CbList(inputCb, lastCb)
	}

	c.callbackQueryQueue = c.callbackQueryQueue[1:]
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
	case "edit":
		c.Edit(inputMsg)
	default:
		return
	}
}

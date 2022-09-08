package commands

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *DummyProductCommander) Get(inputMessage *tgbotapi.Message) {
	productID := strings.Split(inputMessage.CommandArguments(), " ")[0]

	var curMsg string

	found, err := getProductById(&c.service, productID)

	if err != nil {
		curMsg = err.Error()
	} else {
		curMsg = found
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		curMsg,
	)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*DummyProductCommander).Get
}

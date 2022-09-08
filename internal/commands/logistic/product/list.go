package commands

import (
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *DummyProductCommander) List(inputMessage *tgbotapi.Message) {
	limit := strings.Split(inputMessage.CommandArguments(), " ")[0]
	convertedLimit, err := strconv.ParseUint(limit, 10, 64)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid command arg")

		c.bot.Send(msg)

		return
	}

	pLen := c.service.ProductsLength()

	if pLen == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "There is no elements")

		c.bot.Send(msg)

		return
	}

	var pagesQuantity uint64

	if pLen <= convertedLimit {
		pagesQuantity = 1
	} else {
		pagesQuantity = pLen / convertedLimit
	}

	keyboard := tgbotapi.NewInlineKeyboardRow()

	for i := 1; i <= int(pagesQuantity); i++ {
		strPage := strconv.Itoa(i)

		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData(strPage, strPage))
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please, choose the page: ")

	if len(keyboard) != 0 {
		msg.ReplyMarkup = keyboard
	}

	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*DummyProductCommander).List
}

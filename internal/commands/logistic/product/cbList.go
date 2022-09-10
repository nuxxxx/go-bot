package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCbList(c *DummyProductCommander, data string, cq CallbackQueryItem) (string, error) {
	page, err := strconv.ParseUint(data, 10, 64)

	limit := cq.additionalData.(uint64)

	if err != nil {
		return "", err
	}

	products, err := c.service.List(page, limit)

	if err != nil {
		return "", err
	}

	res := ""

	for _, product := range products {
		res += product.String()
	}

	return res, nil
}

func (c *DummyProductCommander) CbList(inputCb *tgbotapi.CallbackQuery, cq CallbackQueryItem) {
	productsString, err := handleCbList(c, inputCb.Data, cq)

	var curMsg string

	if err != nil {
		curMsg = err.Error()
	} else {
		curMsg = productsString
	}

	msg := tgbotapi.NewMessage(inputCb.Message.Chat.ID, curMsg)

	c.bot.Send(msg)
}

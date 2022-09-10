package commands

import (
	"github/nuxxxcake/go-bot/internal/service/logistic/product"
	"math"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func createKeyboard(pLen, convertedLimit uint64) [][]tgbotapi.InlineKeyboardButton {
	var pagesQuantity uint64

	if pLen <= convertedLimit {
		pagesQuantity = 1
	} else {
		pagesQuantity = pLen / convertedLimit
	}

	rowsQuantity := int(math.Ceil(float64(pagesQuantity) / 8))

	keyboardRows := make([][]tgbotapi.InlineKeyboardButton, rowsQuantity)

	for row := 0; row < rowsQuantity; row++ {
		keyboardRow := tgbotapi.NewInlineKeyboardRow()

		for i := row*8 + 1; i <= int(pagesQuantity); i++ {
			strPage := strconv.Itoa(i)

			keyboardRow = append(keyboardRow, tgbotapi.NewInlineKeyboardButtonData(strPage, strPage))
		}

		keyboardRows[row] = keyboardRow
	}

	return keyboardRows
}

func (c *DummyProductCommander) List(inputMessage *tgbotapi.Message) {
	limit := strings.Split(inputMessage.CommandArguments(), " ")[0]
	convertedLimit, err := strconv.ParseUint(limit, 10, 64)

	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid command arg")

		c.bot.Send(msg)

		return
	}

	pLen := product.ProductsLength()

	if pLen == 0 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "There is no elements")

		c.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please, choose the page: ")

	keyboard := createKeyboard(pLen, convertedLimit)

	if len(keyboard) != 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboard...)
	}

	c.bot.Send(msg)
	c.callbackQueryQueue = append(c.callbackQueryQueue, *NewCallbackQueryItem("list", limit))
}

func init() {
	registeredCommands["list"] = (*DummyProductCommander).List
}

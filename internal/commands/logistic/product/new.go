package commands

import (
	"github/nuxxxcake/go-bot/internal/model/logistic"
	"github/nuxxxcake/go-bot/internal/service/logistic/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *DummyProductCommander) New(inputMessage *tgbotapi.Message) {
	productFields, err := productValidation(inputMessage.CommandArguments())

	var curMsg string

	if err != nil {
		curMsg = err.Error()
	} else {
		curMsg = "Created new product."
	}

	product := logistic.Product{
		ID:       product.ProductsLength(),
		Title:    productFields.Title,
		Price:    productFields.Price,
		Quantity: productFields.Quantity,
	}

	c.service.Create(product)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		curMsg,
	)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["new"] = (*DummyProductCommander).New
}

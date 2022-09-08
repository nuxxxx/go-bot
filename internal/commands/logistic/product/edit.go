package commands

import (
	"errors"
	"github/nuxxxcake/go-bot/internal/model/logistic"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleEdit(args []string) (*logistic.Product, error) {
	if len(args) == 0 {
		return nil, errors.New("not enough args")
	}

	productID, err := strconv.ParseUint(args[0], 10, 64)

	if err != nil {
		return nil, errors.New("invalid productID")
	}

	validProduct, err := productValidation(args[1])

	if err != nil {
		return nil, err
	}

	product := logistic.Product{
		ID:       productID,
		Title:    validProduct.Title,
		Quantity: validProduct.Quantity,
		Price:    validProduct.Price,
	}

	return &product, nil
}

func (c *DummyProductCommander) Edit(inputMessage *tgbotapi.Message) {
	args := strings.Split(inputMessage.CommandArguments(), " ")

	var curMsg string

	product, err := handleEdit(args)

	if err != nil {
		curMsg = err.Error()
	} else {
		c.service.Update(product.ID, *product)

		curMsg = "Updated product."
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		curMsg,
	)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["edit"] = (*DummyProductCommander).Edit
}

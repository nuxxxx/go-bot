package commands

import (
	"errors"
	"github/nuxxxcake/go-bot/internal/model/logistic"
	"github/nuxxxcake/go-bot/internal/service/logistic/product"
	"reflect"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleCreatingProduct(ps *product.ProductService, fields []string) error {
	if len(fields) != reflect.TypeOf(logistic.Product{}).NumField()-1 {
		return errors.New("not enough fields")
	}

	// parsing Title
	title := fields[0]

	// parsing price
	price, err := strconv.ParseFloat(fields[1], 64)

	if err != nil {
		return errors.New("invalid price")
	}

	// parsing quantity
	quantity, err := strconv.Atoi(fields[2])

	if err != nil {
		return errors.New("invalid quantity")
	}

	(*ps).Create(logistic.Product{
		ID:       (*ps).ProductsLength() + 1,
		Title:    title,
		Price:    price,
		Quantity: quantity,
	})

	return nil
}

func (c *DummyProductCommander) New(inputMessage *tgbotapi.Message) {
	fields := strings.Split(inputMessage.CommandArguments(), ",")

	err := handleCreatingProduct(&c.service, fields)

	var curMsg string

	if err != nil {
		curMsg = err.Error()
	} else {
		curMsg = "Created new product."
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		curMsg,
	)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["new"] = (*DummyProductCommander).New
}

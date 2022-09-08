package commands

import (
	"github/nuxxxcake/go-bot/internal/service/logistic/product"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleDelete(ps *product.ProductService, productID string) string {
	found, err := getProductById(ps, productID)

	if err != nil {
		return err.Error()
	}

	parsed, err := strconv.ParseUint(found, 10, 64)

	if err != nil {
		return "invalid id"
	}

	(*ps).Remove(parsed)

	return "Deleted."
}

func (c *DummyProductCommander) Delete(inputMsg *tgbotapi.Message) {
	productID := strings.Split(inputMsg.CommandArguments(), " ")[0]

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		handleDelete(&c.service, productID),
	)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*DummyProductCommander).Get
}

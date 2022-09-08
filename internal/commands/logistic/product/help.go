package commands

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func prepareCommands() string {
	res := make([]string, len(registeredCommands))

	for command := range registeredCommands {
		res = append(res, fmt.Sprintf("/%v", command))
	}

	return strings.Join(res, "\n")
}

func (c *DummyProductCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		prepareCommands(),
	)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["help"] = (*DummyProductCommander).Help
}

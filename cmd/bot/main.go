package main

import (
	"log"
	"os"

	commands "github/nuxxxcake/go-bot/internal/commands/logistic/product"
	"github/nuxxxcake/go-bot/internal/service/logistic/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewDummyProductService()
	commander := commands.NewDummyProductCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update.Message)
	}
}

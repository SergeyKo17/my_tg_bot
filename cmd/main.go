package main

import (
	"context"
	"fmt"
	"log"
	"my-bot/internal/config"
	"my-bot/internal/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := config.NewConfig()

	bot, err := tgbotapi.NewBotAPI(config.APIKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			ctx := context.Background()
			fmt.Println(update.Message.Text)

			msgText := handlers.Handlers{}.MainHandler(ctx, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

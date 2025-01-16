package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("tg_bot_token"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s]: %s\n", update.Message.From.UserName, update.Message.Text)

			messageText := update.Message.Text

			switch messageText {
			case isGratitude(messageText):
				stickerId := tgbotapi.FilePath("sticker.webm")
				sticker := tgbotapi.NewSticker(update.Message.Chat.ID, stickerId)
				sticker.ReplyToMessageID = update.Message.MessageID
				bot.Send(sticker)
			case "/start":
				answerText := fmt.Sprintf("Ассаламу алейкум %s! Напиши название города, я скину тебе какая там температура!", update.Message.From.UserName)
				answerMsg := tgbotapi.NewMessage(update.Message.Chat.ID, answerText)
				answerMsg.ReplyToMessageID = update.Message.MessageID
				bot.Send(answerMsg)
			default:
				answerText := GetWeatherStringByCity(messageText)
				answerMsg := tgbotapi.NewMessage(update.Message.Chat.ID, answerText)
				answerMsg.ReplyToMessageID = update.Message.MessageID
				bot.Send(answerMsg)
			}
		}
	}
}

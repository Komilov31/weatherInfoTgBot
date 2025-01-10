package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	myToken := "7553106835:AAE59UjOl5j-wLrvMiybGqKSwDRkC8wKuhY"

	bot, err := tgbotapi.NewBotAPI(myToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			message := update.Message.Text

			if message == "Ты молодец" {
				stickerId := tgbotapi.FilePath("sticker.webm")
				sticker := tgbotapi.NewSticker(update.Message.Chat.ID, stickerId)
				sticker.ReplyToMessageID = update.Message.MessageID
				bot.Send(sticker)
			} else if message == "/start" {
				answer := fmt.Sprintf("Ассаламу алейкум %s! Напиши название города, я скину тебе какая там температура!", update.Message.From.UserName)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else {
				answer := GetTempByCity(message)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}

package main

import (
	"fmt"
	"log"

	"github.com/Komilov31/weatherInfoBot/cmd/logic"
	"github.com/Komilov31/weatherInfoBot/config"
	"github.com/Komilov31/weatherInfoBot/model"
	"github.com/Komilov31/weatherInfoBot/repository"

	// "github.com/Komilov31/weatherInfoBot/user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Envs.TgBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	db := repository.InitStorage()
	store := repository.NewStore(db)

	handler := logic.NewHandler(store)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s]: %s\n", update.Message.From.UserName, update.Message.Text)

			messageText := update.Message.Text

			switch messageText {
			case logic.IsGratitude(messageText):
				stickerId := tgbotapi.FilePath("static/sticker.webm")
				sticker := tgbotapi.NewSticker(update.Message.Chat.ID, stickerId)
				sticker.ReplyToMessageID = update.Message.MessageID
				bot.Send(sticker)
			case "/start":
				answerText := fmt.Sprintf("Ассаламу алейкум %s! Напиши название города, я скину тебе какая там температура!"+model.DefaultMessage,
					update.Message.From.UserName)
				answerMsg := tgbotapi.NewMessage(update.Message.Chat.ID, answerText)
				answerMsg.ReplyToMessageID = update.Message.MessageID
				bot.Send(answerMsg)
			case "/weather":
				temperature := handler.HandleWeatherCommand(update.Message.Chat.UserName)
				answerMsg := tgbotapi.NewMessage(update.Message.Chat.ID, temperature)
				answerMsg.ReplyToMessageID = update.Message.MessageID
				bot.Send(answerMsg)
			case "/setlocation":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Укажите город")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
				nextMessage := <-updates
				answerText := handler.HandleSetLocationCommand(
					nextMessage.Message.Chat.UserName,
					nextMessage.Message.Text,
				)
				answerMsg := tgbotapi.NewMessage(update.Message.Chat.ID, answerText)
				answerMsg.ReplyToMessageID = update.Message.MessageID
				bot.Send(answerMsg)
			default:
				answerText := logic.GetWeatherStringByCity(messageText)
				answerMsg := tgbotapi.NewMessage(update.Message.Chat.ID, answerText)
				answerMsg.ReplyToMessageID = update.Message.MessageID
				bot.Send(answerMsg)
			}
		}
	}
}

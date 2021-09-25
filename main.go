package main

import (
	"errors"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {

	token, err := getBotTokenFromEnvironment()
	if err != nil {
		log.Panic(err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func getBotTokenFromEnvironment() (string, error) {
	log.Println("Attempting to retrieve Telegram bot token from envrionment.")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load from .env file.")
	}

	found, token := len(os.Getenv("TELEGRAM_API_TOKEN")) > 0, os.Getenv("TELEGRAM_API_TOKEN")
	if !found {
		return "", errors.New("Unable to load the token.")
	}

	return token, nil
}

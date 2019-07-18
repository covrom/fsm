package main

import (
	"log"

	"github.com/covrom/fsm"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("ApiKey")
	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		user := update.Message.From.UserName
		chin, ok := listeners.MessageChan(user)
		if !ok {
			userState := NewUserState(bot, update.Message.Chat.ID, chin)
			go fsm.StateMachine(userState.StartState)
		} else {
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					chin <- update.Message
				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					msg.ReplyToMessageID = update.Message.MessageID
					msg.Text = "I don't understand..."
					if _, err := bot.Send(msg); err != nil {
						log.Println(err)
					}
				}
			} else {
				chin <- update.Message
			}
		}
	}

}

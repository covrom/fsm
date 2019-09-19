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
			chin <- update.Message
		}
	}
}

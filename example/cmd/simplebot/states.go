package main

import (
	"github.com/covrom/fsm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	BtnState1 = "State 1"
	BtnReturn = "Return"
)

type UserState struct {
	bot  *tgbotapi.BotAPI
	chin chan *tgbotapi.Message

	ChatID int64
}

func NewUserState(bot *tgbotapi.BotAPI, chatID int64, chin chan *tgbotapi.Message) *UserState {
	s := &UserState{
		bot:    bot,
		chin:   chin,
		ChatID: chatID,
	}
	return s
}

func (s *UserState) StartState() (fsm.StateFunc, error) {
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(BtnState1),
		),
	)
	msg := tgbotapi.NewMessage(s.ChatID, "Main menu")
	msg.ReplyMarkup = buttons
	s.bot.Send(msg)

	for msg := range s.chin {
		if msg.IsCommand() && msg.Command() == "start" {
			// nothing, we are here
			continue
		}
		switch msg.Text {
		case BtnState1:
			return s.State1State, nil
		default:
			m := tgbotapi.NewMessage(s.ChatID, "not implemented")
			s.bot.Send(m)
		}
	}
	return nil, nil
}

func (s *UserState) State1State() (fsm.StateFunc, error) {
	buttons := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(BtnReturn),
		),
	)
	msg := tgbotapi.NewMessage(s.ChatID, "State 1 menu")
	msg.ReplyMarkup = buttons
	s.bot.Send(msg)

	for msg := range s.chin {
		if msg.IsCommand() && msg.Command() == "start" {
			return s.StartState, nil
		}
		switch msg.Text {
		case BtnReturn:
			return s.StartState, nil
		default:
			m := tgbotapi.NewMessage(s.ChatID, "not implemented")
			s.bot.Send(m)
		}
	}
	return nil, nil
}

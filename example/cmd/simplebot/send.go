package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *UserState) Send(msgtext string) {
	msg := tgbotapi.NewMessage(s.ChatID, msgtext)
	s.bot.Send(msg)
}

func (s *UserState) SendAndCloseKeyboard(msgtext string) {
	msg := tgbotapi.NewMessage(s.ChatID, msgtext)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	s.bot.Send(msg)
}

func (s *UserState) SendWithKeyboard1(msgtext string, buttons ...string) {
	msg := tgbotapi.NewMessage(s.ChatID, msgtext)
	keyboard := make([][]tgbotapi.KeyboardButton, 0, len(buttons)/2+1)
	for i := 0; i < len(buttons); i++ {
		keyboard = append(keyboard, []tgbotapi.KeyboardButton{
			tgbotapi.NewKeyboardButton(buttons[i]),
		})
	}
	if len(keyboard) > 0 {
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(keyboard...)
	}
	s.bot.Send(msg)
}

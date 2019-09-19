package main

import (
	"errors"

	"github.com/covrom/fsm"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var NotProcessed = errors.New("Неизвестная команда, попробуйте другую!")
var NoChangeState = errors.New("состояние не меняется")

type ProcessorFunc func() (fsm.StateFunc, error)

type ProcessorFuncMap map[string]ProcessorFunc

type UserInputProcessor struct {
	Commands, Texts         ProcessorFuncMap
	OtherCommand, OtherText func(s string) (fsm.StateFunc, error)
}

func (uip *UserInputProcessor) Process(msg *tgbotapi.Message) (fsm.StateFunc, error) {
	if msg.IsCommand() {
		if uip.Commands != nil {
			f := uip.Commands[msg.Command()]
			if f != nil {
				return f()
			}
		}
		if uip.OtherCommand != nil {
			return uip.OtherCommand(msg.Command())
		}
	} else if !msg.IsCommand() {
		if uip.Texts != nil {
			f := uip.Texts[msg.Text]
			if f != nil {
				return f()
			}
		}
		if uip.OtherText != nil {
			return uip.OtherText(msg.Text)
		}
	}
	return nil, NotProcessed
}

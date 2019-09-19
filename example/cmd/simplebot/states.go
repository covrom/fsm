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
	s.SendWithKeyboard1("Main menu", BtnState1)

	uip := &UserInputProcessor{
		Commands: ProcessorFuncMap{
			"start": func() (fsm.StateFunc, error) {
				return s.StartState, nil
			},
		},
		Texts: ProcessorFuncMap{
			BtnState1: func() (fsm.StateFunc, error) {
				return s.State1State, nil
			},
		},
		OtherCommand: func(cmd string) (fsm.StateFunc, error) {
			s.Send("not implemented")
			return nil, NoChangeState
		},
		OtherText: func(txt string) (fsm.StateFunc, error) {
			s.Send("not implemented")
			return nil, NoChangeState
		},
	}

	for msg := range s.chin {
		rs, err := uip.Process(msg)
		if err != nil {
			if err == NoChangeState {
				continue
			}
			s.SendAndCloseKeyboard(err.Error())
			return s.StartState, nil
		}
		return rs, nil
	}
	return nil, nil
}

func (s *UserState) State1State() (fsm.StateFunc, error) {
	s.SendWithKeyboard1("State 1 menu", BtnReturn)

	uip := &UserInputProcessor{
		Commands: ProcessorFuncMap{
			"start": func() (fsm.StateFunc, error) {
				return s.StartState, nil
			},
		},
		Texts: ProcessorFuncMap{
			BtnReturn: func() (fsm.StateFunc, error) {
				return s.StartState, nil
			},
		},
		OtherCommand: func(cmd string) (fsm.StateFunc, error) {
			s.Send("not implemented")
			return nil, NoChangeState
		},
		OtherText: func(txt string) (fsm.StateFunc, error) {
			s.Send("not implemented")
			return nil, NoChangeState
		},
	}

	for msg := range s.chin {
		rs, err := uip.Process(msg)
		if err != nil {
			if err == NoChangeState {
				continue
			}
			s.SendAndCloseKeyboard(err.Error())
			return s.StartState, nil
		}
		return rs, nil
	}
	return nil, nil
}

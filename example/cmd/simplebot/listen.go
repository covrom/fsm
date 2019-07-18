package main

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var listeners = &Listener{
	users: make(map[string]chan *tgbotapi.Message),
}

type Listener struct {
	sync.RWMutex
	users map[string]chan *tgbotapi.Message
}

func (l Listener) MessageChan(user string) (chan *tgbotapi.Message, bool) {
	l.RLock()
	ch, ok := l.users[user]
	l.RUnlock()
	if ok {
		return ch, true
	}
	l.Lock()
	ch = make(chan *tgbotapi.Message)
	l.users[user] = ch
	l.Unlock()
	return ch, false
}

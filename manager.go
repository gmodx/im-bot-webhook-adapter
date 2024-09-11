package imbotwebhookadapter

import (
	"log/slog"
	"sync"
)

var (
	mgr BotManager
)

func init() {
	mgr = BotManager{
		bots: map[string]Bot{},
	}
}

type BotManager struct {
	bots map[string]Bot
	mu   sync.Mutex
}

func (m *BotManager) Send(msgType MsgType, content string) {
	for name, bot := range m.bots {
		err := bot.Send(msgType, content)
		if err != nil {
			slog.Error("Failed to send message", "name", name, "err", err)
		}
	}
}

func (m *BotManager) Register(name string, bot Bot) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.bots[name] = bot
}

func Send(msgType MsgType, content string) {
	mgr.Send(msgType, content)
}

func Register(name string, bot Bot) {
	mgr.Register(name, bot)
}

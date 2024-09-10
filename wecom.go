package imbotwebhookadapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type WeComBot struct {
	hookKey string
}

type WeComMsg struct {
	Msgtype string `json:"msgtype"`

	Text     WeComMsgContent `json:"text"`
	Markdown WeComMsgContent `json:"markdown"`
}

type WeComMsgContent struct {
	Content string `json:"content"`
}

func NewWeComBot(key string) *WeComBot {
	return &WeComBot{
		hookKey: key,
	}
}

func (w *WeComBot) Send(msgType MsgType, content string) error {
	switch msgType {
	case MsgType_Markdown:
		return w.sendMarkdown(content)
	case MsgType_Text:
		return w.sendText(content)
	}

	return nil
}

func (w *WeComBot) sendMarkdown(content string) error {
	msg := WeComMsg{
		Msgtype: "markdown",
		Markdown: WeComMsgContent{
			Content: content,
		},
	}

	return w.sendCore(msg)
}

func (w *WeComBot) sendText(content string) error {
	msg := WeComMsg{
		Msgtype: "text",
		Text: WeComMsgContent{
			Content: content,
		},
	}

	return w.sendCore(msg)
}

func (w *WeComBot) sendCore(msg interface{}) error {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		slog.Error("Json marshal failed", "err", err.Error())
		return err
	}

	reader := bytes.NewReader(jsonBytes)
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%v", w.hookKey)

	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		slog.Error("Failed to send message to wecom server", "err", err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	_, err = client.Do(request)
	if err != nil {
		slog.Error("Failed to send message to wecom server", "err", err.Error())
		return err
	}

	return nil
}

var _ Bot = &WeComBot{}

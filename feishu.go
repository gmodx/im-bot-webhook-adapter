package imbotwebhookadapter

import (
	"encoding/json"

	"github.com/CatchZeng/feishu/pkg/feishu"
)

type FeiShuBot struct {
	token  string
	secret string
	client *feishu.Client
}

type FeiShuCard struct {
	Elements []FeiShuElement `json:"elements"`
}

type FeiShuElement struct {
	Tag     string     `json:"tag"`
	Text    FeiShuText `json:"text"`
	Content string     `json:"content"`
}

type FeiShuText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

func NewFeiShuBot(token, secret string) *FeiShuBot {
	client := feishu.NewClient(token, secret)

	return &FeiShuBot{
		token:  token,
		secret: secret,
		client: client,
	}
}

func (w *FeiShuBot) Send(msgType MsgType, content string) error {
	switch msgType {
	case MsgType_Markdown:
		return w.sendMarkdown(content)
	case MsgType_Text:
		return w.sendText(content)
	}

	return nil
}

func (w *FeiShuBot) sendMarkdown(content string) error {
	msg := feishu.NewInteractiveMessage()
	msg.MsgType = "interactive"
	card := FeiShuCard{
		Elements: []FeiShuElement{
			{
				Tag:     "markdown",
				Content: content,
			},
		},
	}
	cardBytes, _ := json.Marshal(card)
	msg.Card = string(cardBytes)

	_, _, err := w.client.Send(msg)
	return err
}

func (w *FeiShuBot) sendText(content string) error {
	msg := feishu.NewTextMessage()
	msg.Content.Text = content

	_, _, err := w.client.Send(msg)
	return err
}

var _ Bot = &FeiShuBot{}

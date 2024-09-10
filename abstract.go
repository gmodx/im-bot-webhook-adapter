package imbotwebhookadapter

type MsgType int

const (
	MsgType_Text MsgType = iota
	MsgType_Markdown
)

type Bot interface {
	Send(msgType MsgType, content string) error
}

package imbotwebhookadapter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SendText(t *testing.T) {
	IntegrationTest(t)

	bot := getWeComBot()
	err := bot.sendText("hi")

	assert.NoError(t, err)
}

func Test_SendMarkdown(t *testing.T) {
	IntegrationTest(t)

	bot := getWeComBot()
	err := bot.sendMarkdown(`实时新增用户反馈<font color="warning">132例</font>，请相关同事注意。
> 类型:<font color="comment">用户反馈</font>
> 普通用户反馈:<font color="comment">117例</font>
> VIP用户反馈:<font color="comment">15例</font>`)

	assert.NoError(t, err)
}

func getWeComBot() *WeComBot {
	key := os.Getenv("WECOM_WEBHOOK_KEY")
	bot := NewWeComBot(key)

	return bot
}

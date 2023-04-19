package util

import "github.com/dylanpeng/nbc-chat/lib/chatgpt"

var ChatGPTClient *chatgpt.Client

func InitChatGPTClient(c *chatgpt.Config) (err error) {
	ChatGPTClient, err = chatgpt.NewClient(c)
	return
}

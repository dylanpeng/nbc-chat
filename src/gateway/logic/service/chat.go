package service

import (
	"context"
	"github.com/dylanpeng/nbc-chat/common"
	"github.com/dylanpeng/nbc-chat/common/exception"
	"github.com/dylanpeng/nbc-chat/gateway/util"
	"github.com/sashabaranov/go-openai"
)

var Chat = &chatSrv{}

type chatSrv struct{}

func (s *chatSrv) Chat(ctx context.Context, content string) (rsp openai.ChatCompletionResponse, err *exception.Exception) {
	rsp, e := util.ChatGPTClient.CreateChatCompletion(ctx, openai.ChatMessageRoleUser, content)

	if e != nil {
		err = exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("chatSrv Chat CreateChatCompletion failed. | content: %s | err: %s", content, e)
		return
	}

	return
}

func (s *chatSrv) Completion(ctx context.Context, prompt string) (rsp openai.CompletionResponse, err *exception.Exception) {
	rsp, e := util.ChatGPTClient.CreateCompletion(ctx, prompt)

	if e != nil {
		err = exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("chatSrv Completion CreateCompletion failed. | content: %s | err: %s", prompt, e)
		return
	}

	return
}

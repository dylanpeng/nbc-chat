package control

import "C"
import (
	"github.com/dylanpeng/nbc-chat/common"
	"github.com/dylanpeng/nbc-chat/common/consts"
	ctrl "github.com/dylanpeng/nbc-chat/common/control"
	"github.com/dylanpeng/nbc-chat/gateway/logic/service"
	"github.com/dylanpeng/nbc-chat/lib/proto/chat"
	"github.com/gin-gonic/gin"
)

var Chat = &chatCtrl{}

type chatCtrl struct{}

func (c *chatCtrl) Chat(ctx *gin.Context) {
	req := &chat.ChatReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Content == "") {
		return
	}

	response, err := service.Chat.Chat(ctx, req.Content)

	if err != nil {
		common.Logger.Infof("chatCtrl Chat failed. | content: %s | err: %s", req.Content, err)
		ctrl.Exception(ctx, err)
		return
	}

	rsp := &chat.ChatRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &chat.ChatResult{
			Id:      response.ID,
			Object:  response.Object,
			Created: response.Created,
			Model:   response.Model,
			Usage: &chat.Usage{
				PromptTokens:     int32(response.Usage.PromptTokens),
				CompletionTokens: int32(response.Usage.CompletionTokens),
				TotalTokens:      int32(response.Usage.TotalTokens),
			},
			Message: response.Choices[0].Message.Content,
		},
	}

	ctrl.SendRsp(ctx, rsp)
}

func (c *chatCtrl) Completion(ctx *gin.Context) {
	req := &chat.CompletionReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Prompt == "") {
		return
	}

	response, err := service.Chat.Completion(ctx, req.Prompt)

	if err != nil {
		common.Logger.Infof("chatCtrl Completion failed. | content: %s | err: %s", req.Prompt, err)
		ctrl.Exception(ctx, err)
		return
	}

	rsp := &chat.CompletionRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
		Data: &chat.ChatResult{
			Id:      response.ID,
			Object:  response.Object,
			Created: response.Created,
			Model:   response.Model,
			Usage: &chat.Usage{
				PromptTokens:     int32(response.Usage.PromptTokens),
				CompletionTokens: int32(response.Usage.CompletionTokens),
				TotalTokens:      int32(response.Usage.TotalTokens),
			},
			Message: response.Choices[0].Text,
		},
	}

	ctrl.SendRsp(ctx, rsp)
}

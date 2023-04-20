package control

import (
	"github.com/dylanpeng/nbc-chat/common"
	"github.com/dylanpeng/nbc-chat/common/consts"
	ctrl "github.com/dylanpeng/nbc-chat/common/control"
	"github.com/dylanpeng/nbc-chat/common/exception"
	"github.com/dylanpeng/nbc-chat/gateway/logic/service"
	"github.com/dylanpeng/nbc-chat/gateway/util"
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

func (c *chatCtrl) ListModels(ctx *gin.Context) {
	response, e := util.ChatGPTClient.ListModels(ctx)

	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("ListModels failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	ctrl.SendRsp(ctx, response)
}

func (c *chatCtrl) Edits(ctx *gin.Context) {
	req := &chat.EditsReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Input == "" || req.Instruction == "") {
		return
	}

	modelType := "test"
	if req.ModelType == chat.EditsModelType_CODE {
		modelType = "code"
	}

	response, e := util.ChatGPTClient.Edits(ctx, req.Input, req.Instruction, modelType)

	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("ListModels failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	ctrl.SendRsp(ctx, response)
}

func (c *chatCtrl) CreateImage(ctx *gin.Context) {
	req := &chat.CreateImageReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Prompt == "") {
		return
	}

	response, e := util.ChatGPTClient.CreateImage(ctx, req.Prompt, req.Size, req.Format)

	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("CreateImage failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	if req.Format == consts.ChatImageFormatTypeJson {
		service.Images.SaveImageWithBase64Json("./images", response.Data[0].B64JSON)
	} else {
		service.Images.SaveImageWithUrl("./images", response.Data[0].URL)
	}

	ctrl.SendRsp(ctx, response)
}

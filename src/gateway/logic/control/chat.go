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
	"os"
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
		common.Logger.Infof("ListModels failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
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
		common.Logger.Infof("ListModels failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	ctrl.SendRsp(ctx, response)
}

func (c *chatCtrl) Embedding(ctx *gin.Context) {
	req := &chat.EmbeddingReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Input == "") {
		return
	}

	response, e := util.ChatGPTClient.CreateEmbeddings(ctx, req.Input)

	if e != nil {
		common.Logger.Infof("Embedding CreateEmbeddings failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	ctrl.SendRsp(ctx, response)
}

func (c *chatCtrl) CheckRGBA(ctx *gin.Context) {
	file1, _ := os.Open("./images/image_edit_mask.png")
	file2, _ := os.Open("./images/5525b358-8e33-4aa7-b05c-5d6075790d04.png")

	isTrue1, e1 := common.IsRGBAImage(file1)
	isTrue2, e2 := common.IsRGBAImage(file2)

	common.Logger.Infof("%s, %s, %s, %s", isTrue1, isTrue2, e1, e2)
}

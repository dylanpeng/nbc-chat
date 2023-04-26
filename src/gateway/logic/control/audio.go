package control

import (
	"github.com/dylanpeng/nbc-chat/common"
	ctrl "github.com/dylanpeng/nbc-chat/common/control"
	"github.com/dylanpeng/nbc-chat/common/exception"
	"github.com/dylanpeng/nbc-chat/gateway/config"
	"github.com/dylanpeng/nbc-chat/gateway/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path"
	"strings"
)

var Audio = &audioCtrl{}

type audioCtrl struct{}

func (c *audioCtrl) CreateTranscription(ctx *gin.Context) {
	fileHeader, e := ctx.FormFile("file")
	if e != nil {
		common.Logger.Infof("CreateTranscription FormFile failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	if fileHeader == nil {
		common.Logger.Infof("CreateTranscription file is nil failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeInvalidParams))
		return
	}

	language, _ := ctx.GetPostForm("language")

	fileExt := strings.ToLower(path.Ext(fileHeader.Filename))

	filePath := config.GetConfig().ChatGPT.FileBasePath + "/" + uuid.New().String() + fileExt

	e = ctx.SaveUploadedFile(fileHeader, filePath)

	if e != nil {
		common.Logger.Infof("CreateTranscription SaveUploadedFile failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	response, e := util.ChatGPTClient.CreateTranscription(ctx, filePath, language)

	if e != nil {
		common.Logger.Infof("CreateTranscription CreateTranscription failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	ctrl.SendRsp(ctx, response)
}

func (c *audioCtrl) CreateTranslation(ctx *gin.Context) {
	fileHeader, e := ctx.FormFile("file")
	if e != nil {
		common.Logger.Infof("CreateTranscription FormFile failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	if fileHeader == nil {
		common.Logger.Infof("CreateTranscription file is nil failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeInvalidParams))
		return
	}

	prompt, _ := ctx.GetPostForm("prompt")

	fileExt := strings.ToLower(path.Ext(fileHeader.Filename))

	filePath := config.GetConfig().ChatGPT.FileBasePath + "/" + uuid.New().String() + fileExt

	e = ctx.SaveUploadedFile(fileHeader, filePath)

	if e != nil {
		common.Logger.Infof("CreateTranscription SaveUploadedFile failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	response, e := util.ChatGPTClient.CreateTranslation(ctx, filePath, prompt)

	if e != nil {
		common.Logger.Infof("CreateTranscription CreateTranscription failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	ctrl.SendRsp(ctx, response)
}

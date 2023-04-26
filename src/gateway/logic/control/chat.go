package control

import (
	"bufio"
	"github.com/dylanpeng/nbc-chat/common"
	"github.com/dylanpeng/nbc-chat/common/consts"
	ctrl "github.com/dylanpeng/nbc-chat/common/control"
	"github.com/dylanpeng/nbc-chat/common/exception"
	"github.com/dylanpeng/nbc-chat/gateway/logic/service"
	"github.com/dylanpeng/nbc-chat/gateway/util"
	"github.com/dylanpeng/nbc-chat/lib/proto/chat"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image"
	"image/draw"
	"image/png"
	"mime/multipart"
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

func (c *chatCtrl) EditImage(ctx *gin.Context) {
	prompt, exists := ctx.GetPostForm("prompt")
	if !exists || prompt == "" {
		err := exception.New(exception.CodeInvalidParams)
		common.Logger.Infof("EditImage prompt can't be nil.")
		ctrl.Exception(ctx, err)
		return
	}

	size, _ := ctx.GetPostForm("size")

	fileHeader, exists := ctx.Get(consts.CtxValueImage)
	if !exists {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("EditImage no image failed.")
		ctrl.Exception(ctx, err)
		return
	}

	imageFilePath := "./images/" + uuid.New().String() + ".png"
	ctx.SaveUploadedFile(fileHeader.(*multipart.FileHeader), imageFilePath)

	file, e := fileHeader.(*multipart.FileHeader).Open()
	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("EditImage Open failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	imageFile, e := png.Decode(bufio.NewReader(file))

	m := image.NewRGBA(image.Rect(0, 0, imageFile.Bounds().Max.X/3, imageFile.Bounds().Max.Y/3))
	draw.Draw(imageFile.(draw.Image), image.Rect(imageFile.Bounds().Max.X/3, imageFile.Bounds().Max.Y/3, imageFile.Bounds().Max.X/3*2, imageFile.Bounds().Max.Y/3*2), m, image.ZP, draw.Src)

	maskFilePath := "./images/" + uuid.New().String() + ".png"
	maskFile, e := os.Create(maskFilePath)
	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("EditImage Create failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	e = png.Encode(maskFile, imageFile)
	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("EditImage Encode failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	file.Close()
	maskFile.Close()

	response, e := util.ChatGPTClient.EditImage(ctx, imageFilePath, maskFilePath, prompt, size)

	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("CreateImage failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	if len(response.Data) > 0 && response.Data[0].URL != "" {
		service.Images.SaveImageWithUrl("./images", response.Data[0].URL)
	}

	ctrl.SendRsp(ctx, response)
}

func (c *chatCtrl) VariationImage(ctx *gin.Context) {
	prompt, exists := ctx.GetPostForm("prompt")
	if !exists || prompt == "" {
		err := exception.New(exception.CodeInvalidParams)
		common.Logger.Infof("EditImage prompt can't be nil.")
		ctrl.Exception(ctx, err)
		return
	}

	size, _ := ctx.GetPostForm("size")

	fileHeader, exists := ctx.Get(consts.CtxValueImage)
	if !exists {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("EditImage no image failed.")
		ctrl.Exception(ctx, err)
		return
	}

	file, exists := fileHeader.(*multipart.FileHeader)
	if !exists {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("EditImage Open failed.")
		ctrl.Exception(ctx, err)
		return
	}

	filePath := "./images/" + uuid.New().String() + ".png"
	e := ctx.SaveUploadedFile(file, filePath)
	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("EditImage SaveUploadedFile failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	osFile, e := os.Open(filePath)
	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("EditImage Open failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	defer osFile.Close()

	response, e := util.ChatGPTClient.VariationImage(ctx, osFile, size)

	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("CreateImage failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	if len(response.Data) > 0 && response.Data[0].URL != "" {
		service.Images.SaveImageWithUrl("./images", response.Data[0].URL)
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

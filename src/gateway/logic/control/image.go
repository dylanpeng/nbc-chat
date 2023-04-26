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

var Image = &imageCtrl{}

type imageCtrl struct{}

func (c *imageCtrl) CreateImage(ctx *gin.Context) {
	req := &chat.CreateImageReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Prompt == "") {
		return
	}

	response, e := util.ChatGPTClient.CreateImage(ctx, req.Prompt, req.Size, req.Format)

	if e != nil {
		common.Logger.Infof("CreateImage failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	if req.Format == consts.ChatImageFormatTypeJson {
		service.Images.SaveImageWithBase64Json("./images", response.Data[0].B64JSON)
	} else {
		service.Images.SaveImageWithUrl("./images", response.Data[0].URL)
	}

	ctrl.SendRsp(ctx, response)
}

func (c *imageCtrl) EditImage(ctx *gin.Context) {
	prompt, exists := ctx.GetPostForm("prompt")
	if !exists || prompt == "" {
		common.Logger.Infof("EditImage prompt can't be nil.")
		ctrl.Exception(ctx, exception.New(exception.CodeInvalidParams))
		return
	}

	size, _ := ctx.GetPostForm("size")

	fileHeader, exists := ctx.Get(consts.CtxValueImage)
	if !exists {
		common.Logger.Infof("EditImage no image failed.")
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	imageFilePath := "./images/" + uuid.New().String() + ".png"
	ctx.SaveUploadedFile(fileHeader.(*multipart.FileHeader), imageFilePath)

	file, e := fileHeader.(*multipart.FileHeader).Open()
	if e != nil {
		common.Logger.Infof("EditImage Open failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	defer file.Close()

	imageFile, e := png.Decode(bufio.NewReader(file))

	m := image.NewRGBA(image.Rect(0, 0, imageFile.Bounds().Max.X/2, imageFile.Bounds().Max.Y/2))
	draw.Draw(imageFile.(draw.Image), image.Rect(imageFile.Bounds().Max.X/4, imageFile.Bounds().Max.Y/4, imageFile.Bounds().Max.X/4*3, imageFile.Bounds().Max.Y/4*3), m, image.ZP, draw.Src)

	maskFilePath := "./images/" + uuid.New().String() + ".png"
	maskFile, e := os.Create(maskFilePath)
	if e != nil {
		common.Logger.Infof("EditImage Create failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	defer maskFile.Close()

	e = png.Encode(maskFile, imageFile)
	if e != nil {
		common.Logger.Infof("EditImage Encode failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	response, e := util.ChatGPTClient.EditImage(ctx, imageFilePath, maskFilePath, prompt, size)

	if e != nil {
		common.Logger.Infof("CreateImage failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	if len(response.Data) > 0 && response.Data[0].URL != "" {
		service.Images.SaveImageWithUrl("./images", response.Data[0].URL)
	}

	ctrl.SendRsp(ctx, response)
}

func (c *imageCtrl) VariationImage(ctx *gin.Context) {
	size, _ := ctx.GetPostForm("size")
	fileHeader, exists := ctx.Get(consts.CtxValueImage)

	if !exists {
		common.Logger.Infof("VariationImage no image failed.")
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	imageFilePath := "./images/" + uuid.New().String() + ".png"
	e := ctx.SaveUploadedFile(fileHeader.(*multipart.FileHeader), imageFilePath)

	if e != nil {
		common.Logger.Infof("VariationImage SaveUploadedFile failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	response, e := util.ChatGPTClient.VariationImage(ctx, imageFilePath, size)

	if e != nil {
		common.Logger.Infof("VariationImage VariationImage failed. | err: %s", e)
		ctrl.Exception(ctx, exception.New(exception.CodeQueryFailed))
		return
	}

	if len(response.Data) > 0 && response.Data[0].URL != "" {
		service.Images.SaveImageWithUrl("./images", response.Data[0].URL)
	}

	ctrl.SendRsp(ctx, response)
}

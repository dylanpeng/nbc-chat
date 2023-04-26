package middleware

import (
	"github.com/dylanpeng/golib/coder"
	"github.com/dylanpeng/nbc-chat/common"
	"github.com/dylanpeng/nbc-chat/common/consts"
	ctrl "github.com/dylanpeng/nbc-chat/common/control"
	"github.com/dylanpeng/nbc-chat/common/exception"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path"
	"strings"
)

func JsonCoder(ctx *gin.Context) {
	common.SetCtxCoder(ctx, coder.EncodingJson)
	ctx.Next()
}

func CheckEncoding(ctx *gin.Context) {
	common.SetCtxCoder(ctx, ctx.GetHeader(coder.EncodingHeader))
	ctx.Next()
}

func CrossDomain(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, UserId")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusOK)
	}

	ctx.Next()
}

func Trace(ctx *gin.Context) {
	traceId := ctx.GetHeader(consts.HeaderKeyTraceId)

	if traceId != "" {
		ctx.Set(consts.CtxValueTraceId, traceId)
		return
	}

	trace, exist := ctx.Get(consts.CtxValueTraceId)

	if exist {
		ctx.Set(consts.CtxValueTraceId, trace.(string))
		return
	}

	ctx.Set(consts.CtxValueTraceId, uuid.New().String())

	ctx.Next()
}

func GetImage(ctx *gin.Context) {
	fileHeader, e := ctx.FormFile("file")
	if e != nil {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("GetImage FormFile failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	fileExt := strings.ToLower(path.Ext(fileHeader.Filename))
	if fileExt != ".png" {
		err := exception.New(exception.CodeQueryFailed)
		common.Logger.Infof("GetImage not png failed. | err: %s", e)
		ctrl.Exception(ctx, err)
		return
	}

	if fileHeader.Size > consts.SystemFileSizeLimit {
		err := exception.New(exception.CodeFileSizeOutOfLimit)
		common.Logger.Infof("GetImage file size out of limit failed.")
		ctrl.Exception(ctx, err)
		return
	}

	ctx.Set(consts.CtxValueImage, fileHeader)

	ctx.Next()
}

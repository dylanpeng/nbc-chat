package control

import (
	libHttp "github.com/dylanpeng/golib/http"
	"github.com/dylanpeng/nbc-chat/common/consts"
	ctrl "github.com/dylanpeng/nbc-chat/common/control"
	"github.com/dylanpeng/nbc-chat/common/exception"
	"github.com/dylanpeng/nbc-chat/lib/proto/chat"
	"github.com/gin-gonic/gin"
	"time"
)

var GoogleProxy = &googleProxyCtrl{}

type googleProxyCtrl struct{}

func (c *googleProxyCtrl) Google(ctx *gin.Context) {
	clientProxy := libHttp.NewClientProxy("http://127.0.0.1:58591", 5*time.Second)

	header := make(map[string]string)
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"

	_, body, e := clientProxy.Get("https://www.google.com", header, nil)
	if e != nil {
		err := exception.New(exception.CodeCallHttpFailed, e)
		ctrl.Exception(ctx, err)
		return
	}

	rsp := &chat.CompletionRsp{
		Code:    consts.RespCodeSuccess,
		Message: string(body),
	}

	ctrl.SendRsp(ctx, rsp)
}

func (c *googleProxyCtrl) GoogleNoProxy(ctx *gin.Context) {
	client := libHttp.NewClient(5 * time.Second)
	header := make(map[string]string)
	header["Connection"] = "keep-alive"
	header["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"

	_, body, e := client.Get("https://www.google.com", header, nil)
	if e != nil {
		err := exception.New(exception.CodeCallHttpFailed, e)
		ctrl.Exception(ctx, err)
		return
	}

	rsp := &chat.CompletionRsp{
		Code:    consts.RespCodeSuccess,
		Message: string(body),
	}

	ctrl.SendRsp(ctx, rsp)
}

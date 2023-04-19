package control

import "C"
import (
	"github.com/dylanpeng/nbc-chat/common/consts"
	ctrl "github.com/dylanpeng/nbc-chat/common/control"
	"github.com/dylanpeng/nbc-chat/lib/proto/chat"
	"github.com/gin-gonic/gin"
)

var Chat = &chatCtrl{}

type chatCtrl struct{}

func (c *chatCtrl) Completion(ctx *gin.Context) {
	req := &chat.CompletionReq{}

	if !ctrl.DecodeReq(ctx, req) {
		return
	}

	if !ctrl.ParamAssert(ctx, req, req.Model == "") {
		return
	}

	rsp := &chat.CompletionRsp{
		Code:    consts.RespCodeSuccess,
		Message: consts.RespMsgSuccess,
	}
	ctrl.SendRsp(ctx, rsp)
}

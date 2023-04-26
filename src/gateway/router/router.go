package router

import (
	"github.com/dylanpeng/nbc-chat/common"
	ctrl "github.com/dylanpeng/nbc-chat/common/control"
	"github.com/dylanpeng/nbc-chat/common/middleware"
	"github.com/dylanpeng/nbc-chat/gateway/logic/control"
	"github.com/gin-gonic/gin"
)

var Router = &router{}

type router struct{}

func (r *router) GetIdentifier(ctx *gin.Context) string {
	return common.GetTraceId(ctx)
}

func (r *router) RegHttpHandler(app *gin.Engine) {
	app.Any("/health", ctrl.Health)
	app.Use(middleware.CheckEncoding)
	app.Use(middleware.CrossDomain)
	app.Use(middleware.Trace)

	chatGroup := app.Group("/chatgpt")
	{
		chatGroup.POST("/completion", control.Chat.Completion)
		chatGroup.POST("/chat", control.Chat.Chat)
		chatGroup.GET("/models/list", control.Chat.ListModels)
		chatGroup.POST("/edits", control.Chat.Edits)
		chatGroup.POST("/image/create", control.Chat.CreateImage)
		chatGroup.POST("/image/edit", middleware.GetImage, control.Chat.EditImage)
		chatGroup.POST("/image/variation", middleware.GetImage, control.Chat.VariationImage)
		chatGroup.GET("/google", control.GoogleProxy.Google)
		chatGroup.GET("/google/no-proxy", control.GoogleProxy.GoogleNoProxy)
	}

	testGroup := app.Group("/test")
	{
		testGroup.POST("/rgba", control.Chat.CheckRGBA)
	}
}

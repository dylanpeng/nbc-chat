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
		chatGroup.POST("/embedding", control.Chat.Embedding)

	}

	imageGroup := app.Group("/chatgpt/image")
	{
		imageGroup.POST("/create", control.Image.CreateImage)
		imageGroup.POST("/edit", middleware.GetImage, control.Image.EditImage)
		imageGroup.POST("/variation", middleware.GetImage, control.Image.VariationImage)
	}

	audioGroup := app.Group("/chatgpt/audio")
	{
		audioGroup.POST("/transcription", control.Audio.CreateTranscription)
		audioGroup.POST("/translation", control.Audio.CreateTranslation)
	}

	testGroup := app.Group("/chatgpt/test")
	{
		testGroup.GET("/google", control.GoogleProxy.Google)
		testGroup.GET("/google/no-proxy", control.GoogleProxy.GoogleNoProxy)
		testGroup.POST("/rgba", control.Chat.CheckRGBA)
	}
}

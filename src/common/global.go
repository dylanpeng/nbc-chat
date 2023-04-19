package common

import (
	"github.com/dylanpeng/golib/http"
	"github.com/dylanpeng/golib/logger"
	"github.com/dylanpeng/nbc-chat/common/config"
)

var Logger *logger.Logger
var HttpServer *http.Server

func InitLogger() (err error) {
	conf := config.GetConfig().Log
	Logger, err = logger.NewLogger(conf)
	return err
}

func InitHttpServer(router http.Router) {
	c := config.GetConfig().Server.Http
	HttpServer = http.NewServer(c, router, Logger)
	HttpServer.Start()
}

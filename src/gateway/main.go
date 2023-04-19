package main

import (
	"flag"
	"github.com/dylanpeng/nbc-chat/common"
	"github.com/dylanpeng/nbc-chat/gateway/config"
	"github.com/dylanpeng/nbc-chat/gateway/router"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	configFile = flag.String("c", "config.toml", "config file path")
)

func main() {
	// parse flag
	flag.Parse()

	// set max cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	// parse config file
	if err := config.Init(*configFile); err != nil {
		log.Fatalf("Fatal Error: can't parse config file!!!\n%s", err)
	}

	// init log
	if err := common.InitLogger(); err != nil {
		log.Fatalf("Fatal Error: can't initialize logger!!!\n%s", err)
	}

	defer func() {
		_ = common.Logger.Sync()
		_ = common.Logger.Close()
	}()

	// start http server
	common.InitHttpServer(router.Router)
	common.Logger.Infof("http server start at <%s>", config.GetConfig().Server.Http.GetAddr())

	// waitting for exit signal
	exit := make(chan os.Signal, 1)
	stopSigs := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSigs...)

	// catch exit signal
	sign := <-exit
	common.Logger.Infof("stop by exit signal '%s'", sign)

	// stop http server
	common.HttpServer.Stop()
	common.Logger.Info("http server stoped")

	return
}

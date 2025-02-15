package main

import (
	"danmu-core/internal/manager"
	"danmu-core/internal/model"
	"danmu-core/internal/server"
	"danmu-core/logger"
	"danmu-core/setting"
	"os"
	"os/signal"
	"syscall"
)

func init() {

}

func main() {
	setting.Init()
	logger.Init()
	model.Init()
	manager.InitDouyinManager()
	rpcserver := server.NewRPCServer()
	err := rpcserver.Start()
	if err != nil {
		logger.Fatal().Err(err).Msg("rpc-server start fail")
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 等待退出信号
	<-quit
	logger.Info().Msg("Shutting down server...")

	rpcserver.Stop()
	manager.CloseDouyinManager()
	model.Close()

	logger.Info().Msg("Server exited")
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sxy/demo/config"
	"sxy/demo/pkg/cache"
	"sxy/demo/pkg/db"
	"sxy/demo/pkg/router"
	zaplog "sxy/demo/pkg/zap"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	zaplog.Logger.Debug("START SERVER: ", zap.String("time: ", time.Now().String()))
	config.InitConfig()
	cache.InitRedis()
	db.DbInit()
	defer zaplog.Logger.Sync()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		r := router.NewRouter()
		r.Run(fmt.Sprintf(":%d", config.AppConfig.System.HttpPort))
	}()
	<-c
	zaplog.Logger.Info("shutting down server...")
}

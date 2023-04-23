package main

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	logger.Info("lo-info", zap.String("string", "str"), zap.Int("int", 9))
	sl := logger.Sugar()
	sl.Info("sl-info ")
}

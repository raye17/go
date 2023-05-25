package main

import (
	"go.uber.org/zap"
)

func main() {
	zap.L().Info("global Logger before")
	zap.S().Info("global SugaredLogger before")

	logger := zap.NewExample()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	zap.ReplaceGlobals(logger)
	zap.L().Info("global Logger after")
	zap.S().Info("global SugaredLogger after")
}

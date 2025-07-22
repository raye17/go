package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logFile := "./log/dev.log"
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		if err := os.MkdirAll("./log", os.ModePerm); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}
	writer, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	level := zapcore.DebugLevel
	core := zapcore.NewCore(encoder, zapcore.AddSync(writer), level)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	logger.Info("hello zap")
	logger.Error("error zap")
	logger.Debug("debug zap")
	logger.Warn("warn zap")
}

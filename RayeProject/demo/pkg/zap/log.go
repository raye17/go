package zaplog

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	// 配置日志编码器
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	//encoder := zapcore.NewJSONEncoder(encoderConfig)
	// encoderConfig.MessageKey = "message"
	// encoderConfig.LevelKey = "level"
	// encoderConfig.TimeKey = "time"
	// encoderConfig.CallerKey = "caller"
	// 配置日志级别
	level := zapcore.DebugLevel
	dir := "../log"
	filePath := dir + "/" + time.Now().Format("2006-01-02") + ".log"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	// 创建日志核心
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile)),
		level)

	// 创建日志记录器
	Logger = zap.New(core)

}

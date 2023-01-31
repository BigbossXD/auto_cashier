package utils

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var currentTime = time.Now().Format("02-01-2006")

var logPath = "./logs/applog_" + currentTime + ".log"

var Logger *zap.Logger

func InitializeLogger() {
	c := zap.NewProductionConfig()
	c.EncoderConfig.TimeKey = "timestamp"
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	c.OutputPaths = []string{"stdout", logPath}
	Logger, _ = c.Build()
}

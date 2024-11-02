package logger

import (
	"github.com/dddong3/Bid_Backend/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.SugaredLogger

func init() {
	isProd := config.GetEnv("IS_PROD", "") == "true"
	logLevel := config.GetLogLevel()

	logPath := "bid-backend.log"
	if isProd {
		logPath = "/var/log/bid-backend.log"
	}

	InitLogger(isProd, logPath, logLevel)
}

func InitLogger(isProd bool, logPath string, logLevel string) {
	// var logConfig zap.Config
	// if isProd {
	// 	logConfig = zap.NewProductionConfig()
	// } else {
	// 	logConfig = zap.NewDevelopmentConfig()
	// }
	logConfig := zap.NewDevelopmentConfig()

	logConfig.OutputPaths = []string{logPath, "stdout"}

	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	level := zapcore.InfoLevel
	if err := level.UnmarshalText([]byte(logLevel)); err == nil {
		logConfig.Level.SetLevel(level)
	} else {
		logConfig.Level.SetLevel(zapcore.InfoLevel)
	}

	zapLogger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	Logger = zapLogger.Sugar()
}

func Sync() {
	if Logger != nil {
		err := Logger.Sync()
		if err != nil && err != os.ErrInvalid {
			Logger.Warnw("Failed to sync logger", "error", err)
		}
	}
}

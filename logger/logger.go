package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.SugaredLogger

// InitLogger 初始化日誌系統
func InitLogger(isProd bool, logPath string, logLevel string) {
	// var logConfig zap.Config
	// if isProd {
	// 	logConfig = zap.NewProductionConfig()
	// } else {
	// 	logConfig = zap.NewDevelopmentConfig()
	// }
	logConfig := zap.NewDevelopmentConfig()

	// 動態設置日誌輸出路徑
	logConfig.OutputPaths = []string{logPath, "stdout"}

	// 設置時間格式
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 設置日誌級別
	level := zapcore.InfoLevel // 默認為 info 級別
	if err := level.UnmarshalText([]byte(logLevel)); err == nil {
		logConfig.Level.SetLevel(level)
	} else {
		logConfig.Level.SetLevel(zapcore.InfoLevel) // 如果無效，使用默認的 info 級別
	}

	// 構建日誌器
	zapLogger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	Logger = zapLogger.Sugar()
}

// Sync 清空緩存的日誌
func Sync() {
	if Logger != nil {
		err := Logger.Sync()
		if err != nil && err != os.ErrInvalid {
			// 處理 Sync 時可能發生的錯誤
			Logger.Warnw("Failed to sync logger", "error", err)
		}
	}
}
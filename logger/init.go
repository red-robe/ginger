package logger

import (
	"github.com/gofuncchan/ginger/config"
	"go.uber.org/zap"
)

var ZapLogger *zap.Logger
var ZapHookLogger *zap.Logger

func Init() {
	ZapLogger = zapLogger()

	// 如果开启mongo异步记录日志，则记录mongo hook的日志信息
	if config.LogConf.LogMongoHookSwitch {
		ZapHookLogger = zapHookLogger()
	}

	ZapLogger.Info("zap logger init successful...",
		zap.String("type", "boot log"),
	)
}

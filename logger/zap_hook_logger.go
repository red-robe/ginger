package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

/*
记录zap hook 异步日志中的日志信息，一般在hook的实现内使用，记录异步信息出现的问提，如mongo记录异常，消息队列记录异常等
该类别的日志信息只输出到std和日志文件
*/
func zapHookLogger() *zap.Logger {
	var appName, version, caller, development zap.Option

	// Option：基本日志选项
	appName = zap.Fields(zap.String("app", config.AppName))
	version = zap.Fields(zap.String("version", config.AppVersion))

	// Option：注释每条信息所在文件名和行号
	caller = zap.AddCaller()
	// Option：进入开发模式，使其具有良好的动态性能,记录死机而不是简单地记录错误。
	if gin.IsDebugging() {
		development = zap.Development()
	}

	// 配置核心
	cores := zapcore.NewTee(getHookCoreList()...)

	return zap.New(cores, appName, version, caller, development)
}

/**
获取为自定义hook提供的core列表: 用于记录自定义hook的日志
*/
func getHookCoreList() (coreList []zapcore.Core) {
	// 异步日志信息只记录error级别以上的
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	HookLog := getRotatelogsHook(config.LogConf.LogDir + "hook_error.log")

	// 构建hook的WriteSyncer列表
	var writeSyncerList []zapcore.WriteSyncer
	writeSyncerList = append(writeSyncerList, zapcore.AddSync(os.Stdout), zapcore.AddSync(HookLog))

	coreList = append(coreList,
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(encoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(writeSyncerList...),
			// 日志级别
			errorLevel,
		))

	return
}

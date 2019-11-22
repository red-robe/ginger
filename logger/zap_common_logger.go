package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

/*
项目中使用的通用zap日志记录器，返回logger
*/
func zapLogger() *zap.Logger {
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
	cores := zapcore.NewTee(getCoreList()...)

	return zap.New(cores, appName, version, caller, development)
}

/**
获取zap core列表
*/
func getCoreList() (coreList []zapcore.Core) {
	// 按实际需求灵活定义日志级别
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.WarnLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.ErrorLevel
	})

	fatalLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.FatalLevel
	})

	// 构建hook的 WriteSyncer 列表
	var infoWriteSyncerList, warnWriteSyncerList, errorWriteSyncerList, fatalWriteSyncerList []zapcore.WriteSyncer

	if config.LogConf.LogMongoHookOn {
		// 如开启日志写入mongodb ，将mongo hook 加入WriteSyncer
		mongoLogHook := &MongoLogHook{}
		infoWriteSyncerList = append(infoWriteSyncerList, zapcore.AddSync(mongoLogHook))
		warnWriteSyncerList = append(warnWriteSyncerList, zapcore.AddSync(mongoLogHook))
		errorWriteSyncerList = append(errorWriteSyncerList, zapcore.AddSync(mongoLogHook))
		fatalWriteSyncerList = append(fatalWriteSyncerList, zapcore.AddSync(mongoLogHook))
	} else {
		// 默认输出到文件和std
		rotatelogsInfoHook := getRotatelogsHook(config.LogConf.LogDir + "info.log")
		rotatelogsWarnHook := getRotatelogsHook(config.LogConf.LogDir + "warn.log")
		rotatelogsErrorHook := getRotatelogsHook(config.LogConf.LogDir + "error.log")
		rotatelogsFailHook := getRotatelogsHook(config.LogConf.LogDir + "fail.log")

		infoWriteSyncerList = append(infoWriteSyncerList, zapcore.AddSync(os.Stdout), zapcore.AddSync(rotatelogsInfoHook))
		warnWriteSyncerList = append(warnWriteSyncerList, zapcore.AddSync(os.Stdout), zapcore.AddSync(rotatelogsWarnHook))
		errorWriteSyncerList = append(errorWriteSyncerList, zapcore.AddSync(os.Stdout), zapcore.AddSync(rotatelogsErrorHook))
		fatalWriteSyncerList = append(fatalWriteSyncerList, zapcore.AddSync(os.Stdout), zapcore.AddSync(rotatelogsFailHook))
	}

	coreList = append(coreList,
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(encoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(infoWriteSyncerList...),
			// 日志级别
			infoLevel,
		),
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(encoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(warnWriteSyncerList...),
			// 日志级别
			warnLevel,
		),
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(encoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(errorWriteSyncerList...),
			// 日志级别
			errorLevel,
		),
		zapcore.NewCore(
			// 编码器配置
			zapcore.NewJSONEncoder(encoderConfig()),
			// 打印到控制台和文件
			zapcore.NewMultiWriteSyncer(fatalWriteSyncerList...),
			// 日志级别
			fatalLevel,
		),
	)

	return
}

/*
获取rotatelogs Hook
*/
func getRotatelogsHook(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	rotateLogHook, err := rotatelogs.New(
		filename+".%Y-%m-%d %H:%M:%S",
		rotatelogs.WithLinkName(filename),
		// 最多保留多久
		rotatelogs.WithMaxAge(time.Hour*time.Duration(config.LogConf.LogMaxDayCount*24)),
		// 多久做一次归档
		rotatelogs.WithRotationTime(time.Hour*24*time.Duration(config.LogConf.WithRotationTime)),
	)

	if err != nil {
		panic(err)
	}
	return rotateLogHook
}

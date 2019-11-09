package logger

import (
	"go.uber.org/zap"
	"runtime"
	"strconv"
)

/*
通用业务日志记录工具函数，用于业务日志的记录，区别于框架的中间件请求日志，可写入console，文件，或数据库
*/


// 信息日志记录：通知用户一个操作或状态发生了变化，只作为正常输出；
// 由业务逻辑直接调用
func InfoLog(where,info string)  {
	var msg string
	ZapLog.Info(msg,
		zap.String("where", where),
		zap.String("info", info),
	)
}

// 警告日志记录：当前或未来潜在问题（比如响应速度慢、连接断开、内存吃紧等等)；
// 由业务逻辑直接调用
func WarmLog(err error)  {
	var path string
	if _, file, line, ok := runtime.Caller(1);ok {
		path = file + "" + strconv.Itoa(line)
	}
	msg := "[Warming In "+ path +"]"
	ZapLog.Warn(msg,
		zap.String("path", path),
		zap.String("warming", err.Error()),
	)
}


// 通用错误处理日志记录：错误已经发生，特别区分handler和model位置的错误；
// 由错误处理函数调用，path在错误处理跟踪
func ErrorLog(t, path string, err error) {
	var msg string
	switch t {
	case "Eh":
		msg = "[Error In Handler]"
	case "Em":
		msg = "[Error In Model]"
	case "Ec":
		msg = "[Error In "+ path +"]"
	}

	ZapLog.Error(msg,
		zap.String("type",t),
		zap.String("path", path),
		zap.String("error", err.Error()),
	)
}

// 退出程序的日志记录：致命错误导致程序退出；
// 由Ef函数调用，path在错误处理跟踪
func FailLog(path string, err error) {
	msg := "[Fail In "+ path +"]"
	ZapLog.Fatal(msg,
		zap.String("path", path),
		zap.String("fail", err.Error()),
	)

}


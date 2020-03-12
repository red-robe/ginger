package logger

import (
	"io"
	"log"
	"runtime"
	"strconv"
)

var SimpleLogger *log.Logger

func Init(out io.Writer) {
	if out != nil {
		SimpleLogger = log.New(out, "\t「 ", log.LstdFlags)
	}
}

type consoleColorModeValue int

const (
	autoColor consoleColorModeValue = iota
)

var (
	green            = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white            = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow           = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red              = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue             = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta          = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan             = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset            = string([]byte{27, 91, 48, 109})
	consoleColorMode = autoColor
)

/*
通用业务日志记录工具函数，用于业务日志的记录，区别于框架的中间件请求日志，可写入console，文件，或数据库
*/

// 信息日志记录：通知用户一个操作或状态发生了变化，只作为正常输出；
// 由业务逻辑直接调用
func InfoLog(info string) {
	var path string
	if _, file, line, ok := runtime.Caller(1); ok {
		path = file + "" + strconv.Itoa(line)
	}
	SimpleLogger.Printf("%s [INFO] %s | WHERE: %s | MESSAGE: %s",
		cyan,reset,
		path,
		info,
	)
}

// 警告日志记录：当前或未来潜在问题（比如响应速度慢、连接断开、内存吃紧等等)；
// 由业务逻辑直接调用
func WarmLog(info string) {
	var path string
	if _, file, line, ok := runtime.Caller(1); ok {
		path = file + "" + strconv.Itoa(line)
	}
	SimpleLogger.Printf("%s [WARMING] %s | WHERE: %s | MESSAGE: %s",
		yellow,reset,
		path,
		info,
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
		msg = "[Error In " + path + "]"
	}

	SimpleLogger.Printf("%s [ERROR] %s | WHERE: %s | MESSAGE: %s",
		red,reset,
		msg,
		err.Error(),
	)

}

// 退出程序的日志记录：致命错误导致程序退出；
// 由Ef函数调用，path在错误处理跟踪
func FailLog(path string, err error) {

	msg := "[Fail In " + path + "]"
	SimpleLogger.Printf("%s [FAIL] %s | WHERE: %s | MESSAGE: %s",
		magenta,reset,
		msg,
		err.Error(),
	)
}

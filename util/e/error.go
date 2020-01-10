package e

import (
	"github.com/gofuncchan/ginger/util/logger"
	"runtime"
	"strconv"
)

// handler的错误处理函数：日志记录错误信息
func Eh(err error) bool {
	if err != nil {
		var path = "handler"
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + "" + strconv.Itoa(line)
		}
		logger.ErrorLog("Eh", path, err)
		return false
	}
	return true
}

// model层的错误处理函数：日志记录错误信息
func Em(err error) bool {
	if err != nil {
		var path = "model"
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + "" + strconv.Itoa(line)
		}
		logger.ErrorLog("Em", path, err)
		return false
	}
	return true
}

// 通用的错误处理日志记录
func Ec(err error) bool {
	if err != nil {
		var path = "common"
		if _, file, line, ok := runtime.Caller(1); ok {
			path = file + "" + strconv.Itoa(line)
		}
		logger.ErrorLog("Ec", path, err)
		return false
	}
	return true
}



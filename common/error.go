package common

import "log"

// handler的错误处理函数：日志记录错误信息
func Eh(where string, err error) bool {
	if err != nil {
		log.Println("[Handler Error]:On ->", where, "  [Error Content] ->", err.Error())
		return false
	}
	return true
}

// 模型层的错误处理函数：日志记录错误
func Em(where string, err error) bool {
	if err != nil {
		log.Println("[Model Error]:On ->", where, "  [Error Content] ->", err.Error())
		return false
	}
	return true
}

// 通用的错误处理日志记录
func Ec(where string, err error) bool {
	if err != nil {
		log.Fatal("[Common Error]:On ->", where, "  [Error Content] ->", err.Error())
		return false
	}
	return true
}

// 退出程序的错误处理
func Ef(where string, err error) {
	if err != nil {
		log.Fatal("[Fatal]:On ->", where, "  [Error Content] ->", err.Error())
	}
}

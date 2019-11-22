package common

import (
	"fmt"
	"os"
)

// 退出程序的错误处理,一般用于初始化失败时的错误输出，此时程序正在启动时发生错误，std打印即可。
func EF(err error) {
	if err != nil {
		fmt.Println("Fatal:", err.Error())
		os.Exit(-1)
	}
}

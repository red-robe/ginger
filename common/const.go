package common

import "github.com/gin-gonic/gin"

const (
	// 环境变量未设置时的默认值，默认使用开发环境配置
	DefaultEnv = gin.DebugMode
	// jwt编码时的私钥字符串
	TokenPrivateKey = "ginger"

)

package config

import "github.com/gin-gonic/gin"

// 常量级参数配置项，编译前设置值
const (
	AppName    = "ginger"
	AppVersion = "v0.2.0"
	// 环境变量未设置时的默认值，默认使用开发环境配置
	DefaultEnv = gin.DebugMode
	// jwt编码时的私钥字符串
	TokenPrivateKey = "ginger"
)

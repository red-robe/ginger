package config

import "github.com/gin-gonic/gin"

/*
常量级参数配置项，编译前设置值
*/
const (

	// 基本设置
	AppName    = "ginger"
	AppVersion = "v0.2.0"
	DefaultEnv = gin.DebugMode // 环境变量未设置时的默认值，默认使用开发环境配置

	// jwt编码时的私钥字符串
	TokenPrivateKey = "ginger"

	// 消息中间件的topic

)

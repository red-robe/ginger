package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/boot"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/config"
	ginger_zap_logger "github.com/gofuncchan/ginger/middleware/logger"
	ginger_zap_recovery "github.com/gofuncchan/ginger/middleware/recovery"
	"github.com/gofuncchan/ginger/router"
	"github.com/gofuncchan/ginger/util/logger"
	"strconv"
)

func main() {
	var err error

	// 系统模块初始化
	boot.Init()

	// 创建一个gin实例
	engine := gin.New()

	// zap 日志库
	// zapLogger, _ := zap.NewProduction() //使用默认生存环境配置
	// zapLogger, _ := zap.NewDevelopment() //使用默认开发环境配置
	zapLogger := logger.ZapLog // 使用自定义配置
	defer zapLogger.Sync() // 刷新所有缓冲的日志条目。
	engine.Use(ginger_zap_logger.GingerWithZap(zapLogger))
	engine.Use(ginger_zap_recovery.GingerRecoveryWithZap(zapLogger, true))

	// 路由设置
	router.Router(engine)

	// 设置页面模板路径
	engine.LoadHTMLGlob("views/*")

	err = engine.Run(":" + strconv.Itoa(int(config.BaseConf.ListenPort)))
	common.Ef(err)


}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gofuncchan/ginger/boot"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/config"
	// "github.com/gofuncchan/ginger/cron"
	// "github.com/gofuncchan/ginger/logger"
	"github.com/gofuncchan/ginger/middleware/cors"
	// ginger_zap_logger "github.com/gofuncchan/ginger/middleware/logger"
	// ginger_zap_recovery "github.com/gofuncchan/ginger/middleware/recovery"
	"github.com/gofuncchan/ginger/router"
	// "github.com/gofuncchan/ginger/subscriber/natsSub"
	// "github.com/gofuncchan/ginger/subscriber/redisSub"
	"strconv"
)

func main() {
	// 系统模块初始化
	boot.Init()

	var err error

	// // redis subscriber 运行
	// go redisSub.Run()
	//
	// // nats subscriber 运行
	// go natsSub.Run()
	//
	// // cron 运行
	// go cron.Run()

	// 创建一个gin实例
	engine := gin.Default()

	// engine := gin.New()
	// zap 日志库
	// zapLogger, _ := zap.NewProduction() //使用默认生存环境配置
	// zapLogger, _ := zap.NewDevelopment() //使用默认开发环境配置
	// // 使用自定义配置的zap logger
	// defer logger.ZapLogger.Sync() // 退出前刷新所有缓冲的日志
	// engine.Use(ginger_zap_logger.GingerWithZap(logger.ZapLogger))
	// engine.Use(ginger_zap_recovery.GingerRecoveryWithZap(logger.ZapLogger, true))

	engine.Use(cors.CORSMiddleware())

	// 路由设置
	router.Router(engine)

	// 设置页面模板路径
	engine.LoadHTMLGlob("views/*")

	err = engine.Run(":" + strconv.Itoa(int(config.BaseConf.ListenPort)))
	common.EF(err)

}

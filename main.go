package ginger

import (
	"ginger/common"
	"ginger/init"
	"ginger/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 系统模块初始化
	init.Init()

	// 创建一个gin实例
	engine := gin.New()

	// 加载全局中间件
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// 路由设置
	router.Router(engine)

	// 设置页面模板路径
	engine.LoadHTMLGlob("view/*")

	err := engine.Run(":8090")
	common.Ef("engine.Run",err)
}

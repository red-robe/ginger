package init

import (
	"ginger/config"
	"ginger/dao/mongodb"
	"ginger/dao/mysql"
	"ginger/dao/redis"
	"ginger/util"
)

// 系统启动时的运行各种初始化
func Init()  {
	// 初始化配置参数
	config.Init()
	// 初始化Mysql连接池
	mysql.Init()
	// 初始化Redis连接池
	redis.Init()
	// 初始化mongodb连接池
	mongodb.Init()
	// jwt初始化
	util.JwtInit()
}
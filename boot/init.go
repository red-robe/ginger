package boot

import (
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/dao/mysql"
	"github.com/gofuncchan/ginger/dao/redis"
	"github.com/gofuncchan/ginger/util/jwt"
	"github.com/gofuncchan/ginger/util/logger"
)

// 系统启动时的运行各种初始化
func Init()  {
	// 初始化配置参数
	config.Init()
	// 初始化zap日志
	logger.Init()
	// 初始化Mysql连接池
	mysql.Init()
	// 初始化Redis连接池
	redis.Init()
	// 初始化mongodb连接池
	// mongodb.Init()
	// jwt初始化
	jwt.JwtInit()
}
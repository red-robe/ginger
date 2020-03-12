package boot

import (
	"github.com/gofuncchan/ginger/dao/mongodb"
	"github.com/gofuncchan/ginger/dao/mysql"
	"github.com/gofuncchan/ginger/dao/redis"
)

func CloseDao()  {
	mongodb.CloseSession()
	mysql.CloseDB()
	redis.CloseRedisPool()
}
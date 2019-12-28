package tokenCache

import (
	redigo "github.com/garyburd/redigo/redis"
	"github.com/gofuncchan/ginger/dao/redis"
	"github.com/gofuncchan/ginger/util/e"
)

const UserTokenCacheKeyPrefix = "user.token."

// 设置Token到redis
func SetToken(key, token string) bool {
	rs, err := redigo.String(redis.R("SET", key, token, "EX", 3600))
	if !e.Ec(err){
		return false
	}
	return rs == "OK"
}

// 从redis获取token
func GetToken(key string) string {
	rs, err := redigo.String(redis.R("GET", key))
	if !e.Ec(err){
		return ""
	}
	return rs
}

// 从redis删除token
func DeleteToken(key string) int {
	delCount, err := redigo.Int(redis.R("DEL", key))
	if !e.Ec(err){
		return -1
	}
	return delCount
}

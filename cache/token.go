package cache

import (
	"fmt"
	"ginger/dao/redis"
	redigo "github.com/garyburd/redigo/redis"
)

// 设置Token到redis
func SetToken(key, token string) bool {
	rs, _ := redigo.String(redis.R("SET", key, token, "EX", 3600))
	return rs == "OK"
}

// 从redis获取token
func GetToken(key string) string {
	rs, _ := redigo.String(redis.R("GET", key))
	return rs
}

// 从redis删除token
func DeleteToken(key string) int {
	delCount, err := redigo.Int(redis.R("DEL", key))
	if err != nil {
		fmt.Println("Delete Redis Key Error:", err.Error())
	}
	return delCount
}

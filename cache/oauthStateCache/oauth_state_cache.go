package oauthStateCache

import (
	redigo "github.com/garyburd/redigo/redis"
	"github.com/gofuncchan/ginger/dao/redis"
	"github.com/gofuncchan/ginger/util/e"
)

const WechatOauthStateKeyPrefix = "wechat.oauth.state."
const QQOauthStateKeyPrefix = "qq.oauth.state."
const WeibochatOauthStateKeyPrefix = "weibo.oauth.state."

// 设置Token到redis
func SetOauthSate(key, state string) bool {
	rs, err := redigo.String(redis.R("SET", key, state, "EX", 60))
	if !e.Ec(err) {
		return false
	}
	return rs == "OK"
}

// 从redis获取token
func GetOauthSate(key string) string {
	rs, err := redigo.String(redis.R("GET", key))
	if !e.Ec(err) {
		return ""
	}
	return rs
}

// 从redis删除token
func DeleteOauthSate(key string) int {
	delCount, err := redigo.Int(redis.R("DEL", key))
	if !e.Ec(err) {
		return -1
	}
	return delCount
}

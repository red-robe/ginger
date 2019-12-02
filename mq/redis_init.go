package mq

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/util/logger"
	"strconv"
	"time"
)

/*
RedisPubSub用于实时性较高的消息推送，并不保证可靠,实现实时快速的消息异步分发功能。
使用场景：轻量级，高并发，延迟敏感，即时数据分析、秒杀计数器、缓存等，如果传输的数据量大不建议使用redis pubsub
Tips：原则上用于缓存的redis机器与用于pubsub的redis机器分开较好，如实在用同一个，只需在config配置填写一样即可。
*/
var redisMqPoolPtr *redigo.Pool


func redisMqInit() {
	// 配置并获得一个连接池对象的指针
	redisMqPoolPtr = &redigo.Pool{
		// 最大活动链接数。0为无限
		MaxActive: int(config.MqConf.RedisMq.MaxActive),
		// 最大闲置链接数，0为无限
		MaxIdle: int(config.MqConf.RedisMq.MaxIdle),
		// 闲置链接超时时间
		IdleTimeout: time.Duration(config.MqConf.RedisMq.IdleTimeout) * time.Second,
		// 连接池的连接拨号
		Dial: func() (redigo.Conn, error) {
			// 连接
			redisAddr := config.MqConf.RedisMq.DbHost + ":" + strconv.Itoa(int(config.MqConf.RedisMq.DbPort))
			conn, err := redigo.Dial("tcp", redisAddr)
			if err != nil {
				fmt.Println("redis dial fatal:", err.Error())
				return nil, err
			}
			// 权限认证
			if config.MqConf.RedisMq.DbAuth {
				if _, err := conn.Do("Auth", config.MqConf.RedisMq.DbPasswd); err != nil {
					fmt.Println("redis auth fatal:", err.Error())
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},

		// 定时检测连接是否可用
		TestOnBorrow: func(conn redigo.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("Ping")
			if err != nil {
				logger.WarmLog("Redis PubSub Server Disconnect")
			}
			return err
		},
	}

	// 一般启动后不关闭连接池
	// defer poolPtr.Close()

	RedisMq = new(RedisPubSub)
	fmt.Println("Redis PubSub init ready!")

	RedisQueue = GetRedisList(GetRedisConn())
	fmt.Println("Redis List Queue init ready!")

}

// 从Redis连接池获取一个连接
func GetRedisConn() redigo.Conn {
	conn := redisMqPoolPtr.Get()
	return conn
}

// 从Redis连接池获取一个PubSub连接
func GetRedisPubSubConn() redigo.PubSubConn {
	conn := redisMqPoolPtr.Get()
	psConn := redigo.PubSubConn{Conn: conn}
	return psConn
}

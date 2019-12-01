package mq

import (
	"context"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/util/logger"
	"strconv"
	"time"
)

var redisMqPoolPtr *redigo.Pool

// 原则上用于缓存的redis机器与用于pubsub的redis机器分开较好，如实在用同一个，只需在config配置填写一样即可
func redisMqInit()  {
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

	GingerMq = new(RedisMq)
	fmt.Println("Redis PubSub init ready!")


}

// 从Redis连接池获取一个连接
func GetRedisConn() redigo.Conn {
	conn := redisMqPoolPtr.Get()
	return conn
}

func GetRedisPubSubConn() redigo.PubSubConn {
	conn := redisMqPoolPtr.Get()
	psConn := redigo.PubSubConn{Conn: conn}
	return psConn
}

// 实现GingerMQ接口的entity
type RedisMq struct {}

/*
发布消息
@param topic string 发布频道
@param msg interface{} 发布的消息
*/
func (mq *RedisMq) Public(topic string, msg interface{}) error {
	conn := GetRedisConn()
	defer conn.Close()

	receiveNum, err := redigo.Int(conn.Do("PUBLIC", topic, redigo.Args{}.AddFlat(msg)))
	if err != nil {
		return err
	}

	if receiveNum == 0 {
		// 订阅并接收到该topic的数量为receiveNum
		logger.WarmLog(fmt.Sprintf("Nobody subscribe or receive %s topic  ", topic))
	}

	return nil
}

/*
订阅消息处理
@param receiveMsgFunc func(channel string, data []byte) error 接收消息的处理函数
@param channels []string 订阅的频道，允许多个
*/
func (mq *RedisMq) Subscribe(receiveMsgFunc recSubMsgFunc,channels ...string) error {
	psConn := GetRedisPubSubConn()
	defer psConn.Close()

	if err := psConn.Subscribe(redigo.Args{}.AddFlat(channels)); err != nil {
		return err
	}

	// 接收订阅消息处理
	ctx, cancel := context.WithCancel(context.TODO())
	err := receiveRedisSubcribeMessage(
		ctx,
		psConn,
		receiveMsgFunc,
		func() error {
			cancel()
			return nil
		},
	)

	return err
}

// 按模式匹配字符串的频道订阅
func PSubscribe(receiveMsgFunc recSubMsgFunc,channelsWithPattern ...string) error {
	psConn := GetRedisPubSubConn()
	defer psConn.Close()

	if err := psConn.PSubscribe(redigo.Args{}.AddFlat(channelsWithPattern)); err != nil {
		return err
	}

	// 接收订阅消息处理
	ctx, cancel := context.WithCancel(context.TODO())
	err := receiveRedisSubcribeMessage(
		ctx,
		psConn,
		receiveMsgFunc,
		func() error {
			cancel()
			return nil
		},
	)

	return err
}

// 订阅后接收消息
func receiveRedisSubcribeMessage(
	ctx context.Context,
	psConn redigo.PubSubConn,
	onMessage recSubMsgFunc,
	onZeroSub func() error,

) error {

	var err error
	// 启动一个协程去获取订阅topic的返回信息
	done := make(chan error, 1)
	go func() {
		for {
			// 针对不同接收消息处理
			switch n := psConn.Receive().(type) {
			case error:
				// 订阅返回错误信息时发送到done通道
				done <- n
				return
			case redigo.Message:
				// 接收到一般消息时由onMessage函数处理
				if err = onMessage(n.Channel, n.Data); err != nil {
					done <- err
					return
				}
			case redigo.Subscription:
				// 接收到订阅或取消订阅的消息
				if n.Count == 0 {
					// 当前连接订阅数为0时，从goroutine返回。
					err := onZeroSub()
					done <- err
				}
			}
		}
	}()

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
loop:
	for err == nil {
		select {
		case <-ticker.C:
			// 每分钟一次连接的健康监测
			if err := psConn.Ping(""); err != nil {
				break loop
			}
		case <-ctx.Done():
			// 上下文cancel，直接跳出loop
			break loop
		case err = <-done:
			// 接收到done的error信息则直接返回
			return err
		}
	}

	// Wait for goroutine to complete.
	return <-done
}

// 取消订阅
func (mq *RedisMq) Unsubscribe(channels... string) error{
	psConn := GetRedisPubSubConn()
	defer psConn.Close()

	err := psConn.Unsubscribe(redigo.Args{}.AddFlat(channels))

	return err
}

// 按模式匹配字符串的频道取消订阅
func (mq *RedisMq) PUnsubscribe(channelsWithPattern... string) error{
	psConn := GetRedisPubSubConn()
	defer psConn.Close()

	err := psConn.PUnsubscribe(redigo.Args{}.AddFlat(channelsWithPattern))

	return err
}

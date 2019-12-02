package mq

import (
	"context"
	"fmt"
	"github.com/gofuncchan/ginger/util/logger"
	redigo "github.com/garyburd/redigo/redis"

	"time"
)

var RedisMq *RedisPubSub


// 实现GingerMQ接口的entity
type RedisPubSub struct{}

/*
发布消息
@param channel string 发布频道
@param msg interface{} 发布的消息
*/
func (mq *RedisPubSub) Public(channel string, msg interface{}) error {
	conn := GetRedisConn()
	defer conn.Close()

	receiveNum, err := redigo.Int(conn.Do("PUBLIC", channel, redigo.Args{}.AddFlat(msg)))
	if err != nil {
		return err
	}

	if receiveNum == 0 {
		// 订阅并接收到该channel的数量为receiveNum
		logger.WarmLog(fmt.Sprintf("Nobody subscribe or receive %s channel  ", channel))
	}

	return nil
}

/*
订阅消息处理
@param receiveMsgFunc func(channel string, data []byte) error 接收消息的处理函数
@param channels []string 订阅的频道，允许多个
*/
func (mq *RedisPubSub) Subscribe(receiveMsgFunc recSubMsgFunc, channels ...string) error {
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
func  (mq *RedisPubSub) PSubscribe(receiveMsgFunc recSubMsgFunc, channelsWithPattern ...string) error {
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
	// 启动一个协程去获取订阅channel的返回信息
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
func (mq *RedisPubSub) Unsubscribe(channels ...string) error {
	psConn := GetRedisPubSubConn()
	defer psConn.Close()

	err := psConn.Unsubscribe(redigo.Args{}.AddFlat(channels))

	return err
}

// 按模式匹配字符串的频道取消订阅
func (mq *RedisPubSub) PUnsubscribe(channelsWithPattern ...string) error {
	psConn := GetRedisPubSubConn()
	defer psConn.Close()

	err := psConn.PUnsubscribe(redigo.Args{}.AddFlat(channelsWithPattern))

	return err
}

package redisSub

import (
	"fmt"
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/mq/redisMq"
	"github.com/gofuncchan/ginger/util/logger"
)

func Run() {
	// 订阅需开启redismq服务
	if !config.MqConf.RedisMq.Switch {
		return
	}

	fmt.Println("The RedisMq Subscriber Running")
	var err error

	// 整合所有的订阅消息处理函数给订阅器接收消息
	// 演示订阅两个topic的接收消息处理
	var recMsgFuncs = make(map[string]redisMq.RecSubMsgFunc)
	recMsgFuncs[redisMq.RedisTestchannel1] = func(topic string, msg interface{}) error {
		logger.InfoLog("Subscriber receive", fmt.Sprintf("Redis Topic: %s,Message:%s", topic, msg))
		return nil
	}
	recMsgFuncs[redisMq.RedisTestchannel2] = func(topic string, msg interface{}) error {
		logger.InfoLog("Subscriber receive", fmt.Sprintf("Redis Topic: %s,Message:%s", topic, msg))
		return nil
	}

	// 订阅频道，可多个,并接收处理消息
	err = redisMq.Subscribe(recMsgFuncs, redisMq.RedisTestchannel1, redisMq.RedisTestchannel2)
	if err != nil {
		logger.WarmLog(err.Error())
	}

}

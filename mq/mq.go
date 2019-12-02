package mq

import (
	"github.com/gofuncchan/ginger/config"
)

// PubSub 接口
type GingerPubSub interface {
	Public(topic string,msg interface{}) error
	Subscribe(f recSubMsgFunc,topic... string) error
	Unsubscribe(topic... string) error
}
// 订阅模式下的消息处理函数类型
type recSubMsgFunc func(topic string,msg interface{}) error


// Queue 接口
type GingerQueue interface {
	Push(topic string,msg interface{}) error
	Pop(topic string) (interface{},error)
}


func Init()  {
	if config.MqConf.RedisMq.Switch {
		redisMqInit()
	}

	if config.MqConf.NatsMq.Switch {
		natsMqInit()
	}

}

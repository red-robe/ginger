package mq

import (
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/mq/natsMq"
	"github.com/gofuncchan/ginger/mq/redisMq"
)

// 由于消息系统各有特色及用途，暂不做通用化的接口处理
// // PubSub 接口
// type GingerPubSub interface {
// 	Public(topic string,msg interface{}) error
// 	Subscribe(f RecSubMsgFunc,topic... string) error
// 	Unsubscribe(topic... string) error
// }
//
// // Queue 接口
// type GingerQueue interface {
// 	Push(topic string,msg interface{}) error
// 	Pop(topic string) (interface{},error)
// }


func Init()  {
	if config.MqConf.RedisMq.Switch {
		redisMq.RedisMqInit()
	}

	if config.MqConf.NatsMq.Switch {
		natsMq.NatsMqInit()
	}

}

package mq

import (
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/util/logger"
	"github.com/pkg/errors"
)

// MQ 接口
type GingerMQ interface {
	Public(topic string,msg interface{}) error
	Subscribe(f recSubMsgFunc,topic... string) error
	Unsubscribe(topic... string) error
}

// 订阅模式下的消息处理函数类型
type recSubMsgFunc func(topic string,msg interface{}) error

// 用于全局的MQ调用
var GingerMq GingerMQ

func Init()  {
	// 是否打开mq开关
	if !config.MqConf.MqSwitch {
		return
	}

	// 选择一个消息中间件的服务redis|nats|kafka
	switch config.MqConf.MqSelect {
	case "redis":
		redisMqInit()
	case "nats":
		natsMqInit()
	case "kafka":
		kafkaMqInit()
	default:
		logger.FailLog("mq.Init",errors.New("MqConf.MqSelect error: "+config.MqConf.MqSelect))
		return
	}

}

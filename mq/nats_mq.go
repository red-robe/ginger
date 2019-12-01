package mq

import (
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/util/logger"
	"github.com/nats-io/nats.go"
	"time"
)

var NatsConn *nats.EncodedConn

func natsMqInit() {
	nc, err := nats.Connect(config.MqConf.NatsMq.HostUrl)
	if err != nil {
		logger.FailLog("mq.natsMqInit Connect", err)
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		logger.FailLog("mq.natsMqInit NewEncodedConn", err)
	}
	defer c.Close()

	NatsConn = c

	GingerMq = new(NatsMQ)
}

type NatsMQ struct{}

func (mq *NatsMQ) Public(topic string, msg interface{}) error {
	err := NatsConn.Publish(topic, msg)
	return err
}

func (mq *NatsMQ) Subscribe(f recSubMsgFunc, topic ...string) error {
	var handler = func( m *nats.Msg) {
		err := f(m.Subject, m)
		if err != nil {
			logger.WarmLog("recSubMsgFunc Error:" + err.Error())
		}
	}
	for _, t := range topic {
		sub, err := NatsConn.Subscribe(t, handler)
		if err != nil {
			return err
		}

		const MAX_WANTED = 10
		err = sub.AutoUnsubscribe(MAX_WANTED)
		if err != nil {
			return err
		}
	}

	return nil
}


func (mq *NatsMQ) Unsubscribe(topic ...string) error {
	for _, t := range topic {
		sub, err := NatsConn.Subscribe(t, nil)
		if err != nil {
			return err
		}

		err = sub.Unsubscribe()
		if err != nil {
			return err
		}
	}
	return nil
}

// 要求对回复主题做出响应,使用Request()自动内联等待响应
func (mq *NatsMQ) PublicRequest(topic, reply string, msg interface{}) error {
	err := NatsConn.PublishRequest(topic, reply, msg)
	return err
}

// 请求一个主题回复后，使用一个指针变量接收响应信息
func (mq *NatsMQ) Request(topic string, msg interface{}, respPtr interface{}, timeout time.Duration) error {
	err := NatsConn.Request(topic, msg, &respPtr, 10*time.Millisecond)
	if err != nil {
		return err
	}
	return err
}

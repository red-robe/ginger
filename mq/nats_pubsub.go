package mq

import "github.com/nats-io/nats.go"


/*
发布消息到一个主题
@param subject string  发布主题
@param msg interface{} 发布的消息
*/
func (mq *NatsMQ) Public(subject string, msg interface{}) error {
	err := mq.conn.Publish(subject, msg)
	return err
}

/*
订阅并异步接收主题数据
@param subject string  订阅主题
@param cb nats.Handler 订阅消息处理函数
For example：
handler := func(m *Msg)
handler := func(p *person)
handler := func(subject string, o *obj)
handler := func(subject, reply string, o *obj)   for Request() reply
*/
func (mq *NatsMQ) Subscribe(subject string, handler nats.Handler) error {

	sub, err := mq.conn.Subscribe(subject, handler)
	if err != nil {
		return err
	}

	const MAX_WANTED = 10
	err = sub.AutoUnsubscribe(MAX_WANTED)
	if err != nil {
		return err
	}

	return nil
}

/*
取消订阅一个或多个主题
param subject/subjects string 已订阅的主题
*/
func (mq *NatsMQ) Unsubscribe(subjects ...string) error {
	for _, subject := range subjects {
		sub, err := mq.conn.Subscribe(subject, nil)
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

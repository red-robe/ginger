package mq

func kafkaMqInit()  {

}

type KafkaMq struct {}

func (mq *KafkaMq) Public(topic string, msg interface{}) error {
	panic("implement me")
}

func (mq *KafkaMq) Subscribe(f recSubMsgFunc, topic ...string) {
	panic("implement me")
}

func (mq *KafkaMq) Unsubscribe(topic ...string) {
	panic("implement me")
}



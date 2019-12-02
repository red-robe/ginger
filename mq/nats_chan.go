package mq

// 发送消息到一个主题，绑定管道
func (mq *NatsMQ)BindSendChan(subject string,sendCh chan interface{}) error {
	err := mq.conn.BindSendChan(subject, sendCh)
	return err
}

// 接收主题消息，绑定管道
func (mq *NatsMQ)BindRecvChan(subject string,recvCh chan interface{}) error {
	_, err := mq.conn.BindRecvChan(subject, recvCh)
	return err
}

// 基于队列的接收操作，绑定通道。
func (mq *NatsMQ)BindRecvQueueChan(subject,queue string,recvCh chan interface{}) error {
	_, err := mq.conn.BindRecvQueueChan(subject,queue, recvCh)
	return err
}
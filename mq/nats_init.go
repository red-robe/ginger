package mq

import (
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/util/logger"
	"github.com/nats-io/nats.go"
	"strconv"
	"strings"
	"time"
)

/*
NatsMq，类似于redis式的轻量级消息中间件，用于高吞吐量的应用，性能比redis高许多，但不保证可靠送达，消息发送后不管
特性：高性能（fast）、一直可用（dial tone）、极度轻量级（small footprint）、最多交付一次（fire and forget，消息发送后不管）、支持多种消息通信模型和用例场景（flexible）
应用场景：　寻址、发现、命令和控制（控制面板）、负载均衡、多路可伸缩能力、定位透明、容错等。
*/

var NatsMq *NatsMQ

func natsMqInit() {
	var serverList []string
	for _, server := range config.MqConf.NatsMq.NatsServers {
		var natsUrl = "nats://"
		if server.Host == "" {
			natsUrl = nats.DefaultURL
		} else {
			if server.AuthSwitch {
				natsUrl += server.UserName + ":" + server.Password + "@"
			}
			natsUrl += server.Host + ":" + strconv.Itoa(server.Port)
		}
		serverList = append(serverList, natsUrl)
	}

	// TODO nats client 实现连接池
	// 连接nats server
	NatsServersUrl := strings.Join(serverList, ",")
	nc, err := nats.Connect(NatsServersUrl,
		nats.MaxReconnects(5),         // 设置重新连接等待和最大重新连接尝试次数
		nats.ReconnectWait(2*time.Second), // 每次重连等待时间
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			logger.WarmLog("Nats server disconnected Reason:" + err.Error())
		}),                                // 断开连接的错误处理
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.WarmLog("Nats server reconnected to " + nc.ConnectedUrl())
		}),                                // 重连时的错误处理
		nats.ClosedHandler(func(nc *nats.Conn) {
			logger.WarmLog("Nats server connection closed. Reason: " + nc.LastError().Error())
		}),                                // 关闭连接时的错误处理
	)
	if err != nil {
		logger.FailLog("mq.natsMqInit Connect", err)
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		logger.FailLog("mq.natsMqInit NewEncodedConn", err)
	}
	defer c.Close()

	NatsMq = new(NatsMQ)
	NatsMq.conn = c
}

type NatsMQ struct {
	conn *nats.EncodedConn
}

// 收到执行关闭
func (mq *NatsMQ) Close() {
	mq.conn.Close()
}

/*
Flush 当执行完整个服务并接收到所有内部reply时返回
*/
func (mq *NatsMQ) Flush() error {
	return mq.conn.Flush()
}

/*
Flush的超时限制的实现
*/
func (mq *NatsMQ) FlushTimeout(timeout time.Duration) error {
	return mq.conn.FlushTimeout(timeout)
}

/*
Drain将使连接处于排空状态。所有订阅都将立即进入耗尽状态。完成后，发布服务器将被耗尽，并且不能发布任何其他消息。
一旦排空发布服务器，连接将关闭。使用ClosedCB()选项可以知道连接何时已从排出状态移到关闭状态。
*/
func (mq *NatsMQ) Drain() error {
	return mq.conn.Drain()
}
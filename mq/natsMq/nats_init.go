package natsMq

import (
	"fmt"
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

var NatsMqConnPool *NatsPool

func NatsMqInit() {
	var err error
	var serverList []string
	var natsServersUrl string
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
	if len(serverList) > 1 {
		natsServersUrl = strings.Join(serverList, ",")
	} else {
		natsServersUrl = serverList[0]
	}
	fmt.Println("natsServersUrl:",natsServersUrl)


	//  nats conn 初始化连接池
	NatsMqConnPool, err = NewDefaultPool(natsServersUrl)
	if err != nil {
		logger.FailLog("NewDefaultPool Error", err)
	}

}


/*
Flush 当执行完整个服务并接收到所有内部reply时返回
*/
func Flush(conn nats.Conn) error {
	return conn.Flush()
}

/*
Flush的超时限制的实现
*/
func FlushTimeout(conn nats.Conn,timeout time.Duration) error {
	return conn.FlushTimeout(timeout)
}

/*
Drain将使连接处于排空状态。所有订阅都将立即进入耗尽状态。完成后，发布服务器将被耗尽，并且不能发布任何其他消息。
一旦排空发布服务器，连接将关闭。使用ClosedCB()选项可以知道连接何时已从排出状态移到关闭状态。
*/
func Drain(conn nats.Conn) error {
	return conn.Drain()
}
package mongodb

import (
	"fmt"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/config"
	"gopkg.in/mgo.v2"
	"strconv"
	"strings"

	// "gopkg.in/mgo.v2/bson"
)

/*
mongo连接池实现：
由于mgo连接池是自带的，你只需要使用session.Copy()拿到一个复制的session，用完之后session.Close()即可。
*/

var (
	session *mgo.Session
	err     error
)

// 包初始化时实例化一个mongo session
func Init() {
	// 启动一个Session并关闭返回的副本，只为了初始化
	Session()
	fmt.Println("Mongodb session clone init ready!")
}

// copy session 实现连接池
func Session() *mgo.Session {
	if session == nil {
		dbHosts := config.MongoConf.DbHosts
		dbPorts := config.MongoConf.DbPorts
		dbUser := config.MongoConf.DbUser
		dbPassword := config.MongoConf.DbPasswd
		dbName := config.MongoConf.DbName

		var mongoUrl string
		mongoUrl = `mongodb://`

		if dbUser != "" && dbPassword != "" {
			mongoUrl += dbUser + ":" + dbPassword + "@"
		}

		if len(dbHosts) > 1 {
			var hps []string
			for k, v := range dbHosts {
				hp := v + ":" + strconv.Itoa(dbPorts[k])
				hps = append(hps, hp)
			}
			mongoUrl += strings.Join(hps, ",")
		} else if len(dbHosts) == 1 {
			mongoUrl += dbHosts[0] + ":" + strconv.Itoa(dbPorts[0])
		}

		if dbName != "" {
			mongoUrl += "/" + dbName
		}

		session, err = mgo.Dial(mongoUrl)
		if err != nil {
			common.EF(err)
		}
	}
	return session.Clone()
}

package mongodb

import (
	"fmt"
	"github.com/gofuncchan/ginger/common"
	"gopkg.in/mgo.v2"
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
		session, err = mgo.Dial("localhost")
		if err != nil {
			common.EF(err)
		}
	}
	return session.Clone()
}

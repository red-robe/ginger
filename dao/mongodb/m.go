package mongodb

import (
	"github.com/gofuncchan/ginger/config"
	"gopkg.in/mgo.v2"
)

/*
使用示例
	定义一个接收器
	oneData := make(map[string]interface{})

	每执行一个M函数都从连接池拷贝一个连接，闭包内有集合collection的操作句柄实例，直接在闭包里mongo操作
	M("student", func(collection *mgo.Collection) {
		err := collection.FindId(bson.ObjectIdHex("5c8608a7baa5203fc198859f")).One(&oneData)
		ErrorHandler(err, "Find().One()")
		fmt.Println("查询一个结果：", oneData)
	})
*/
func M(collection string, f func(*mgo.Collection)) {
	// 申请一个mongodb连接拷贝
	session := Session()
	// 使用完即释放连接
	defer func() {
		session.Close()
	}()

	// 返回一个collection连接闭包
	c := session.DB(config.MongoConf.DbName).C(collection)
	f(c)
}



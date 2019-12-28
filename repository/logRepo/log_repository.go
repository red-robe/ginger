package logRepo

import (
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/dao/mongodb"
	"gopkg.in/mgo.v2"
	"time"
)

const MongoLogCollection = "log"

func InsertMongoLog(dataMap map[string]interface{}) (err error) {

	// 直接使用M方法记录
	err = mongodb.M(MongoLogCollection, func(c *mgo.Collection) error {
		// mongo日志过期自动清理
		dateTimeFieldName := "dateTime"
		dataMap[dateTimeFieldName] = time.Now()
		datetimeIndexUsedForExpire := mgo.Index{
			Key:         []string{dateTimeFieldName},
			ExpireAfter: time.Duration(config.LogConf.LogMongoExpireAfterSeconds) * time.Second,
		}
		if err = c.EnsureIndex(datetimeIndexUsedForExpire); err != nil {
			return err
		}

		err = c.Insert(dataMap)
		if err != nil {
			return err
		}
		return nil
	})

	return
}

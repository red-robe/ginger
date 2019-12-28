package logger

import (
	"github.com/gofuncchan/ginger/repository/logRepo"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

// 实现WriteSyncer接口的 MongoLogHook，用于异步记录日志到mongo
type MongoLogHook struct {
}

/**
实现MongoLogHook的io.writer接口方法, 使当前对象可以作为zap的hook使用
*/
func (m *MongoLogHook) Write(data []byte) (n int, err error) {
	if err = m.insertLogToMongo(data); err != nil {
		return 0, err
	}

	return
}

/**
插入日志到mongo
*/
func (m *MongoLogHook) insertLogToMongo(data []byte) error {
	// 解析数据为json格式
	var object interface{}
	if err := jsoniter.Unmarshal(data, &object); err != nil {
		ZapHookLogger.Error("jsoniter Unmarshal error", zap.Any("data", data), zap.Error(err))
		return err
	}

	// 转为map类型
	dataMap := object.(map[string]interface{})

	if err := logRepo.InsertMongoLog(dataMap); err != nil {
		ZapHookLogger.Error("insert mongo log error", zap.Any("data", dataMap), zap.Error(err))
		return err
	}

	// logger.ZapHookLogger.Info("insert mongo log successful")

	return nil
}

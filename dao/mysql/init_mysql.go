package mysql

import (
	"database/sql"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/config"
	"github.com/gofuncchan/ginger/util/logger"
	"time"
)

// 基于gendry 管理的sql builder的连接
var (
	Db  *sql.DB
	err error
)

func Init() {
	Db, err = manager.New(config.MysqlConf.DbName, config.MysqlConf.DbUser, config.MysqlConf.DbPasswd, config.MysqlConf.DbHost).Set(
		manager.SetCharset(config.MysqlConf.ChartSet),                                   // 设置编码类型：utf8
		manager.SetAllowCleartextPasswords(config.MysqlConf.AllowCleartextPasswords),    // 开发环境中设置允许明文密码通信
		manager.SetInterpolateParams(config.MysqlConf.InterpolateParams),                // 设置允许占位符参数
		manager.SetTimeout(time.Duration(config.MysqlConf.Timeout)*time.Second),         // 连接超时时间
		manager.SetReadTimeout(time.Duration(config.MysqlConf.ReadTimeout)*time.Second), // 读超时时间
		manager.SetParseTime(config.MysqlConf.ParseTime),                                // 将数据库的datetime时间格式转换为go time包数据类型
	).Port(int(config.MysqlConf.DbPort)).Open(config.MysqlConf.PING)
	common.Ef(err)

	Db.SetConnMaxLifetime(time.Duration(config.MysqlConf.ConnMaxLifetime) * time.Second) // 设置最大的连接时间，1分钟
	Db.SetMaxIdleConns(int(config.MysqlConf.MaxIdleConns))                               // 设置最大的闲置连接数
	Db.SetMaxOpenConns(int(config.MysqlConf.MaxOpenConns))                               // 设置最大的连接数

	err := Db.Ping()
	if err != nil {
		logger.WarmLog(err)
	} else {
		logger.InfoLog("dao/mysql.Init", "Mysql Pool Ready!")
	}
}

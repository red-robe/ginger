package config

import (
	"ginger/common"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Base struct {
	AppName string `yaml:"appName"`
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
	ListenPort int64 `yaml:"listen"`
}

type Mysql struct {
	DbHost                  string `yaml:"dbHost"`
	DbPort                  int64  `yaml:"dbPort"`
	DbUser                  string `yaml:"dbUser"`
	DbPasswd                string `yaml:"dbPasswd"`
	DbName                  string `yaml:"dbName"`
	ConnMaxLifetime         int64  `yaml:"connMaxLifetime"`
	MaxIdleConns            int64  `yaml:"maxIdleConns"`
	MaxOpenConns            int64  `yaml:"maxOpenConns"`
	ChartSet                string `yaml:"charset"`
	AllowCleartextPasswords bool   `yaml:"allowCleartextPasswords"`
	InterpolateParams       bool   `yaml:"interpolateParams"`
	Timeout                 int64  `yaml:"timeout"`
	ReadTimeout             int64  `yaml:"readTimeout"`
	ParseTime               bool   `yaml:"parseTime"`
	PING                    bool   `yaml:"ping"`
}

type Redis struct {
	DbHost      string `yaml:"dbHost"`
	DbPort      int64  `yaml:"dbPort"`
	DbAuth      bool   `yaml:"dbAuth"`
	DbPasswd    string `yaml:"dbPasswd"`
	MaxActive   int64  `yaml:"maxActive"`
	MaxIdle     int64  `yaml:"maxIdle"`
	IdleTimeout int64  `yaml:"idleTimeout"`
}

type Mongodb struct {
	DbHost   string `yaml:"dbHost"`
	DbPort   int64  `yaml:"dbPort"`
	DbUser   string `yaml:"dbUser"`
	DbPasswd string `yaml:"dbPasswd"`
	DbName   string `yaml:"dbName"`
}

var (
	BaseConf  Base
	MysqlConf Mysql
	RedisConf Redis
	MongoConf Mongodb
)

func Init() {
	// 先从环境变量获取环境的信息 (debug|release|test)
	// export GIN_MODE=release
	var currentEnv string
	currentEnv = gin.Mode()
	if currentEnv == "" {
		currentEnv = common.DefaultEnv
	}


	confPath := "./config/" + currentEnv + "/"

	baseConfFile, err := ioutil.ReadFile(confPath + "/base.yaml")
	common.Ef(err)
	err = yaml.Unmarshal(baseConfFile, &BaseConf)
	common.Ef(err)

	mysqlConfFile, err := ioutil.ReadFile(confPath + "/mysql.yaml")
	common.Ef(err)
	err = yaml.Unmarshal(mysqlConfFile, &MysqlConf)
	common.Ef(err)

	RedisConfFile, err := ioutil.ReadFile(confPath + "/redis.yaml")
	common.Ef(err)
	err = yaml.Unmarshal(RedisConfFile, &RedisConf)
	common.Ef(err)

	MongoConfFile, err := ioutil.ReadFile(confPath + "/mongodb.yaml")
	common.Ef(err)
	err = yaml.Unmarshal(MongoConfFile, &MongoConf)
	common.Ef(err)

}

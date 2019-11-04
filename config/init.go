package config

import (
	"ginger/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Base struct {
	AppName string `yaml:"appName"`
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
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
	// 先从环境变量获取环境的信息 (dev|prod|test)
	var currentEnv string
	currentEnv = os.Getenv(common.EnvVariableName)
	if currentEnv == "" {
		currentEnv = common.DefaultProjectEnv
	}

	confPath := "./config/" + currentEnv + "/"

	baseConfFile, err := ioutil.ReadFile(confPath + "/base.yaml")
	common.Ef("ioutil.ReadFile base.yaml", err)
	err = yaml.Unmarshal(baseConfFile, &BaseConf)
	common.Ef("yaml.Unmarshal(baseConfFile, &baseConf)", err)

	mysqlConfFile, err := ioutil.ReadFile(confPath + "/mysql.yaml")
	common.Ef("ioutil.ReadFile mysql.yaml", err)
	err = yaml.Unmarshal(mysqlConfFile, &MysqlConf)
	common.Ef("yaml.Unmarshal(mysqlConfFile, &mysqlConfig)", err)

	RedisConfFile, err := ioutil.ReadFile(confPath + "/redis.yaml")
	common.Ef("ioutil.ReadFile redis.yaml", err)
	err = yaml.Unmarshal(RedisConfFile, &RedisConf)
	common.Ef("yaml.Unmarshal(RedisConfFile, &RedisConf)", err)

	MongoConfFile, err := ioutil.ReadFile(confPath + "/mongodb.yaml")
	common.Ef("ioutil.ReadFile mongodb.yaml", err)
	err = yaml.Unmarshal(MongoConfFile, &MongoConf)
	common.Ef("yaml.Unmarshal(MongoConfFile, &MongoConf)", err)
}

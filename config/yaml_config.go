package config

import (
	"github.com/gofuncchan/ginger/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 基础配置
type Base struct {
	Env        string `yaml:"env"`
	ListenPort int64  `yaml:"listen"`
}


// MysqlDB 配置
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

// RedisDB配置
type Redis struct {
	DbHost      string `yaml:"dbHost"`
	DbPort      int64  `yaml:"dbPort"`
	DbAuth      bool   `yaml:"dbAuth"`
	DbPasswd    string `yaml:"dbPasswd"`
	MaxActive   int64  `yaml:"maxActive"`
	MaxIdle     int64  `yaml:"maxIdle"`
	IdleTimeout int64  `yaml:"idleTimeout"`
}

// MongoDB 配置
type Mongodb struct {
	DbHosts  []string `yaml:"dbHosts"`
	DbPorts  []int    `yaml:"dbPorts"`
	DbUser   string   `yaml:"dbUser"`
	DbPasswd string   `yaml:"dbPasswd"`
	DbName   string   `yaml:"dbName"`
}




// Cors配置
type Cors struct {
	AllowHeaders     []string `yaml:"AllowHeaders"`
	AllowCredentials bool     `yaml:"AllowCredentials"`
	ExposeHeaders    []string `yaml:"ExposeHeaders"`
	MaxAge           int      `yaml:"MaxAge"`
	AllowAllOrigins  bool     `yaml:"AllowAllOrigins"`
	AllowOrigins     []string `yaml:"AllowOrigins"`
	AllowMethods     []string `yaml:"AllowMethods"`
}


var (
	BaseConf  *Base
	MysqlConf *Mysql
	RedisConf *Redis
	MongoConf *Mongodb
	CorsConf  *Cors
)

// 动态参数配置项，编译后可携yaml配置文件启动
func Init(confPath string) {

	baseConfFile, err := ioutil.ReadFile(confPath + "/base.yaml")
	common.EF(err)
	err = yaml.Unmarshal(baseConfFile, &BaseConf)
	common.EF(err)

	corsConfFile, err := ioutil.ReadFile(confPath + "/cors.yaml")
	common.EF(err)
	err = yaml.Unmarshal(corsConfFile, &CorsConf)
	common.EF(err)

	mysqlConfFile, err := ioutil.ReadFile(confPath + "/mysql.yaml")
	common.EF(err)
	err = yaml.Unmarshal(mysqlConfFile, &MysqlConf)
	common.EF(err)

	RedisConfFile, err := ioutil.ReadFile(confPath + "/redis.yaml")
	common.EF(err)
	err = yaml.Unmarshal(RedisConfFile, &RedisConf)
	common.EF(err)

	MongoConfFile, err := ioutil.ReadFile(confPath + "/mongodb.yaml")
	common.EF(err)
	err = yaml.Unmarshal(MongoConfFile, &MongoConf)
	common.EF(err)



}

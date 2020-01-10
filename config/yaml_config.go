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

// 日志配置
type Log struct {
	LogDir                     string `yaml:"logDir"`
	LogMaxDayCount             int    `yaml:"maxDayCount"`
	LogMongoHookSwitch         bool   `yaml:"logMongoHookSwitch"`
	LogMongoCollection         string `yaml:"logMongoCol"`
	LogMongoExpireAfterSeconds int    `yaml:"logMongoExpire"`
	WithRotationTime           int    `yaml:"withRotationTime"`
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

// 消息系统配置
type Mq struct {
	RedisMq `yaml:"redisMq"`
	NatsMq  `yaml:"natsMq"`
}

// Redis Pubsub 消息系统
type RedisMq struct {
	Switch      bool   `yaml:"switch"`
	MaxActive   int    `yaml:"maxActive"`
	MaxIdle     int    `yaml:"maxIdle"`
	IdleTimeout int    `yaml:"idleTimeout"`
	DbHost      string `yaml:"dbHost"`
	DbPort      int    `yaml:"dbPort"`
	DbAuth      bool   `yaml:"dbAuth"`
	DbPasswd    int    `yaml:"dbPasswd"`
}

// Nats Mq 消息系统
type NatsMq struct {
	Switch      bool         `yaml:"switch"`
	NatsServers []NatsServer `yaml:"natsServer"`
}

// 可配集群
type NatsServer struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	AuthSwitch bool   `yaml:"authSwitch"`
	UserName   string `yaml:"userName"`
	Password   string `yaml:"password"`
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

// OSS对象存储配置
type Oss struct {
	Qiniu  `yaml:"qiniu"`
	Aliyun `yaml:"aliyun"`
}

type Qiniu struct {
	Switch           bool   `yaml:"switch"`
	AccessKey        string `yaml:"accessKey"`
	SecretKey        string `yaml:"secretKey"`
	Bucket           string `yaml:"bucket"`
	UseHTTPS         bool   `yaml:"useHTTPS"`
	UseCdnDomains    bool   `yaml:"UseCdnDomains"`
	UpTokenExpires   int    `yaml:"upTokenExpires"`
	CallbackURL      string `yaml:"callbackURL"`
	CallbackBodyType string `yaml:"callbackBodyType"`
	EndUser          string `yaml:"endUser"`
	FsizeMin         int    `yaml:"fsizeMin"`
	FsizeMax         int    `yaml:"fsizeLimit"`
	MimeLimit        string `yaml:"mimeLimit"`
}

type Aliyun struct {
	Switch          bool   `yaml:"switch"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	ConnTimeout     int    `yaml:"connTimeout"`
	RWTimeout       int    `yaml:"rwTimeout"`
	EnableMD5       bool   `yaml:"enableMD5"`
	EnableCRC       bool   `yaml:"enableCRC"`
	AuthProxy       string `yaml:"authProxy"`
	Proxy           string `yaml:"proxy"`
	AccessKeyId     string `yaml:"accessKeyId"`
	BucketName      string `yaml:"bucketName"`
	Endpoint        string `yaml:"endpoint"`
	UseCname        bool   `yaml:"useCname"`
	SecurityToken   string `yaml:"securityToken"`
}

var (
	BaseConf  *Base
	LogConf   *Log
	MysqlConf *Mysql
	RedisConf *Redis
	MongoConf *Mongodb
	MqConf    *Mq
	CorsConf  *Cors
	OssConf   *Oss
)

// 动态参数配置项，编译后可携yaml配置文件启动
func Init(confPath string) {

	baseConfFile, err := ioutil.ReadFile(confPath + "/base.yaml")
	common.EF(err)
	err = yaml.Unmarshal(baseConfFile, &BaseConf)
	common.EF(err)

	logConfFile, err := ioutil.ReadFile(confPath + "/log.yaml")
	common.EF(err)
	err = yaml.Unmarshal(logConfFile, &LogConf)
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

	MqConfFile, err := ioutil.ReadFile(confPath + "/mq.yaml")
	common.EF(err)
	err = yaml.Unmarshal(MqConfFile, &MqConf)
	common.EF(err)

	CorsConfFile, err := ioutil.ReadFile(confPath + "/cors.yaml")
	common.EF(err)
	err = yaml.Unmarshal(CorsConfFile, &CorsConf)
	common.EF(err)

	OssConfFile, err := ioutil.ReadFile(confPath + "/oss.yaml")
	common.EF(err)
	err = yaml.Unmarshal(OssConfFile, &OssConf)
	common.EF(err)

}

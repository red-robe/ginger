package common

const (
	// 系统的环境变量KEY，程序启动时可自动获取，用于选择对应环境的配置文件
	EnvVariableName = "GO_PROJECT_ENV"
	// 环境变量未设置时的默认值，默认使用开发环境配置
	DefaultProjectEnv = "dev"
	// jwt编码时的私钥字符串
	TokenPrivateKey = "gofuncchan"

)

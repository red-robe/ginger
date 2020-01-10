package oss

import (
	aliOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofuncchan/ginger/common"
	"github.com/gofuncchan/ginger/config"
	qiuniuOss "github.com/qiniu/api.v7/v7/auth/qbox"
)

var Qiniu *qiuniuOss.Mac
var Aliyun *aliOss.Client

func Init() {
	// 七牛云存储初始化
	if config.OssConf.Qiniu.Switch {
		Qiniu = qiuniuOss.NewMac(config.OssConf.Qiniu.AccessKey, config.OssConf.Qiniu.SecretKey)
	}

	// Aliyun OSS初始化
	if config.OssConf.Aliyun.Switch {
		var err error
		Aliyun, err = aliOss.New(
			config.OssConf.Aliyun.Endpoint,
			config.OssConf.Aliyun.AccessKeyId,
			config.OssConf.Aliyun.AccessKeySecret,
			aliOss.Timeout(int64(config.OssConf.Aliyun.ConnTimeout), int64(config.OssConf.Aliyun.RWTimeout)),
			aliOss.UseCname(config.OssConf.Aliyun.UseCname),
			aliOss.EnableCRC(config.OssConf.Aliyun.EnableCRC),
		)
		if err != nil {
			common.EF(err)
		}
	}

}

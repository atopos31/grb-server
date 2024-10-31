package service

import (
	"fmt"
	"gvb/config"
	"gvb/models/res"
	"gvb/service/inter"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type OssQinui struct {
	accessKey  string
	secretKey  string
	Domain     string // 域名
	bucket     string // 存储桶
	expire     uint64 // 过期时间
	region     string // 存储区域
	keyPrefix  string // 文件前缀
	imgProcess string // 图片处理
}

func NewOssQinui(config config.OssQiniu) inter.OssService {
	if config.AccessKey == "" || config.SecretKey == "" || config.Bucket == "" || config.Domain == "" || config.Region == "" {
		panic("oss Qiniu config error")
	}
	return &OssQinui{
		accessKey:  config.AccessKey,
		secretKey:  config.SecretKey,
		Domain:     config.Domain,
		bucket:     config.Bucket,
		expire:     config.Expire,
		region:     config.Region,
		keyPrefix:  config.KeyPrefix,
		imgProcess: config.ImgProcess,
	}
}

func (o *OssQinui) GetUploadToken() (res.OssConfig, error) {
	putPolicy := storage.PutPolicy{
		Scope:           fmt.Sprintf("%s:%s", o.bucket, o.keyPrefix),
		Expires:         o.expire,
		IsPrefixalScope: 1,
	}
	// 构建表单上传的对象
	mac := auth.New(o.accessKey, o.secretKey)
	// 生成上传凭证
	uptoken := putPolicy.UploadToken(mac)

	ossConfig := res.OssConfig{
		Domain:     o.Domain,
		Bucket:     o.bucket,
		Region:     o.region,
		KeyPrefix:  o.keyPrefix,
		ImgProcess: o.imgProcess,
		UpToken:    uptoken,
	}
	return ossConfig, nil
}

package service

import (
	"gvb/config"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type OssQinui struct {
	accessKey string
	secretKey string
	bucket    string
}

func NewOssQinui(config config.OssQiniu) *OssQinui {
	return &OssQinui{
		accessKey: config.AccessKey,
		secretKey: config.SecretKey,
		bucket:    config.Bucket,
	}
}

func (o *OssQinui) GetUploadToken() (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: o.bucket,
	}
	mac := auth.New(o.accessKey,o.secretKey)
	uptoken := putPolicy.UploadToken(mac)
	return uptoken, nil
}

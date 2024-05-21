package config

type Oss struct {
	OssQiniu OssQiniu `mapstructure:"qiniu"` // 七牛云OSS配置
}

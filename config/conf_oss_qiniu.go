package config

type OssQiniu struct {
	AccessKey  string `mapstructure:"access_key"`
	SecretKey  string `mapstructure:"secret_key"`
	Domain     string `mapstructure:"domain"`      // 域名
	Bucket     string `mapstructure:"bucket"`      // 存储空间
	Region     string `mapstructure:"region"`      // 存储区域
	Expire     uint64 `mapstructure:"expire"`      // 过期时间
	KeyPrefix  string `mapstructure:"key_prefix"`  // 存储路径前缀
	ImgProcess string `mapstructure:"img_process"` // 图片处理
}

func NewOssQiniuConfig(config Config) OssQiniu {
	return config.Oss.OssQiniu
}

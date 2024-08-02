package config

type AiQianfan struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Model     string `mapstructure:"model"`
}

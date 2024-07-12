package config

type AiHunyuan struct {
	SecretId  string `mapstructure:"secret_id"`
	SecretKey string `mapstructure:"secret_key"`
	Model     string `mapstructure:"model"`
}

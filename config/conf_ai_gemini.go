package config

type AiGemini struct {
	ApiKey string `mapstructure:"api_key" `
	Model  string `mapstructure:"model" `
	Sokcs5 string `mapstructure:"socks5"`
}

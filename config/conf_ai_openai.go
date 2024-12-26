package config

type AiOpenai struct {
	ApiKey  string `mapstructure:"api_key" `
	Model   string `mapstructure:"model" `
	BaseURL string `mapstructure:"base_url"`
}

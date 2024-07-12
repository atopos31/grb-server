package config

type Meilisearch struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	ApiKey string `mapstructure:"api_key"`
}

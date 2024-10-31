package config

import "fmt"

type Meilisearch struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	ApiKey string `mapstructure:"api_key"`
	Index  string `mapstructure:"index"`
}

func NewMeilisearchConfig(config Config) Meilisearch {
	return config.Meilisearch
}

func (m *Meilisearch) Dsn() string {
	return fmt.Sprintf("http://%s:%s", m.Host, m.Port)
}

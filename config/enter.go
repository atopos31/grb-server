package config

type Config struct {
	Sys         System      `mapstructure:"system"`
	Jwt         Jwt         `mapstructure:"jwt"`
	Logger      Logger      `mapstructure:"logger"`
	Mysql       Mysql       `mapstructure:"mysql"`
	Meilisearch Meilisearch `mapstructure:"meilisearch"`
	Redis       Redis       `mapstructure:"redis"`
	Oss         Oss         `mapstructure:"oss"`
	Ai          Ai          `mapstructure:"ai"`
}

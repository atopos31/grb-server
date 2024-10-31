package config

import "github.com/google/wire"

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

var ProvideSet = wire.NewSet(
	NewMeilisearchConfig,
	NewAiConfig,
	NewRedisConfig,
	NewMysqlConfig,
	NewOssQiniuConfig,
	NewSiteInfoPath,
)

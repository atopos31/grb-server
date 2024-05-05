package config

import ()

type Config struct {
	Sys    System `mapstructure:"system"`
	Logger Logger `mapstructure:"logger"`
	Mysql  Mysql  `mapstructure:"mysql"`
}

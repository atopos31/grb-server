package config

type Config struct {
	Sys    System `mapstructure:"system"`
	Logger Logger `mapstructure:"logger"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
}

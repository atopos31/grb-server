package config

type Config struct {
	Sys    System `mapstructure:"system"`
	Jwt    Jwt    `mapstructure:"jwt"`
	Logger Logger `mapstructure:"logger"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
}

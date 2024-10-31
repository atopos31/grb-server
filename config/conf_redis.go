package config

import "fmt"

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func NewRedisConfig(config Config) Redis {
	return config.Redis
}

func (r *Redis) Addr() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

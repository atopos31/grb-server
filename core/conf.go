package core

import (
	"gvb/config"

	"github.com/spf13/viper"
)

func NewConf(configPath string) config.Config {
	confViper := viper.New()

	confViper.SetConfigFile(configPath)
	var config config.Config

	if err := confViper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := confViper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}

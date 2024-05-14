package core

import (
	"time"

	"gvb/config"
	"gvb/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConf(configPath string) *config.Config {
	viper.SetConfigFile(configPath)
	c := new(config.Config)
	var err error
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err = viper.Unmarshal(c); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(c); err != nil {
			global.Log.Error("viper.Unmarshal failed", err)
			return
		}
		global.Log.Debug("config has changed", time.Now().String())
	})

	return c
}

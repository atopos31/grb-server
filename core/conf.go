package core

import (
	"time"

	"gvb/config"
	"gvb/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConf(configPath string) *config.Config {
	confViper := viper.New()

	confViper.SetConfigFile(configPath)
	config := new(config.Config)

	if err := confViper.ReadInConfig(); err != nil {
		println("viper.ReadInConfig failed,please enter " + configPath + "exit!")
		panic(err)
	}

	if err := confViper.Unmarshal(config); err != nil {
		panic(err)
	}

	confViper.WatchConfig()
	confViper.OnConfigChange(func(in fsnotify.Event) {
		if err := confViper.Unmarshal(config); err != nil {
			global.Log.Error("viper.Unmarshal failed", err)
			return
		}
		global.Log.Debug("config has changed", time.Now().String())
	})

	return config
}

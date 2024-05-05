package core

import (
	"fmt"
	"gvb/config"
	"gvb/global"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConf(configPath string) (err error) {
	viper.SetConfigFile(configPath)
	var c = new(config.Config)
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err = viper.Unmarshal(c); err != nil {
		panic(err)
	}
	global.Conf = c

	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config has changed, time:" + time.Now().String())
		if err := viper.Unmarshal(global.Conf); err != nil {
			global.Log.Error("viper.Unmarshal failed", err)
			return
		}
		global.Log.Debug("config has changed")
	})

	return nil
}

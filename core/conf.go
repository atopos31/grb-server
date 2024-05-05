package core

import (
	"fmt"
	"gvb/config"
	"gvb/global"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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
			zap.L().Error("viper umarshal Conf falied,Error:", zap.Error(err))
			return
		}
		zap.L().Debug("config has changed", zap.Any("new config:", global.Conf))
	})

	return nil
}

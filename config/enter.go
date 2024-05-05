package config

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(Config)

type Config struct {
	Sys    System `mapstructure:"system"`
	Logger Logger `mapstructure:"logger"`
	Mysql  Mysql  `mapstructure:"mysql"`
}

type System struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type Logger struct {
	LogLevel     string `mapstructure:"level"` // 日志级别 debugger输出全部SQL dev 输出错误日志 release
	Prefix       string `mapstructure:"prefix"`
	Director     string `mapstructure:"director"`
	ShowLine     bool   `mapstructure:"show_line"`      // 是否显示行号
	LogInConsole bool   `mapstructure:"log_in_console"` // 是否输出到控制台
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	LogLevel string `mapstructure:"log_level"`
}

func Init(configPath string) (err error) {
	viper.SetConfigFile(configPath)
	fmt.Println(configPath)
	if err = viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err = viper.Unmarshal(Conf); err != nil {
		panic(err)
	}
	fmt.Printf("config:%v", Conf)
	// 监听配置文件变化
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config has changed, time:" + time.Now().String())
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper umarshal Conf falied,Error:", zap.Error(err))
			return
		}
		zap.L().Debug("config has changed", zap.Any("new config:", Conf))
	})

	return nil
}

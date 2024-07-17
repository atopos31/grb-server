package main

import (
	"flag"

	"gvb/core"
	"gvb/global"
	"gvb/routers"
	"gvb/service"
)

func main() {
	// 从命令行读取配置文件位置
	configDir := flag.String("config", "./config/dev.yaml", "config dir")
	flag.Parse()

	config := core.NewConf(*configDir)
	global.Conf = &config
	global.Log = core.NewLogger(config.Logger)

	service.Svc = service.New(config)

	router := routers.NewRouter(config.Sys)

	if config.Sys.Env == "debug" {
		global.Log.Infof("[Config]:%v", config)
	}
	router.Run(config.Sys.Addr())
}

package main

import (
	"flag"

	"gvb/app"
	"gvb/config"
	"gvb/core"
	"gvb/routers"
	"gvb/service"
)

func main() {
	// 从命令行读取配置文件位置
	configDir := flag.String("config", "./config/dev.yaml", "config dir")
	flag.Parse()

	conf := core.NewConf(*configDir)

	app.Conf = &conf
	app.Log = core.NewLogger(conf.Logger)

	service.Svc = service.New(conf)

	server := routers.New(conf.Sys)

	if conf.Sys.Env == config.ENV_SYS_DEBUG {
		app.Log.Infof("[Config]:%v", conf)
		app.PrintLogo()
	}

	server.Run(conf.Sys.Addr())
}

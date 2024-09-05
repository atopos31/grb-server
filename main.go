package main

import (
	"flag"

	"gvb/app"
	"gvb/core"
	"gvb/routers"
	"gvb/service"
)

func main() {
	// 从命令行读取配置文件位置
	configPath := flag.String("config", "./config/dev.yaml", "config dir")
	flag.Parse()
	conf := core.NewConf(*configPath)
	app.Conf = &conf
	app.Log = core.NewLogger(conf.Logger)
	service.Svc = service.New(conf)
	server := routers.New(conf.Sys)

	app.Run(server)
}

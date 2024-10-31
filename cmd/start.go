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
	configPath := flag.String("config", "../config/dev.yaml", "config dir")
	flag.Parse()
	conf := core.NewConf(*configPath)
	app.Conf = &conf
	app.Log = core.NewLogger(conf.Logger)
	svc, err := service.New(conf)
	if err != nil {
		panic(err)
	}
	service.Svc = svc
	server := routers.New(conf.Sys)

	app.Run(server)
}

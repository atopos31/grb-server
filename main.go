package main

import (
	"flag"
	"gvb/core"
	"gvb/global"
	"gvb/routers"
)

func main() {
	configDir := flag.String("config", "./config/dev.yaml", "config dir")
	flag.Parse()
	// 初始化配置
	core.InitConf(*configDir)
	// 初始化日志
	global.Log = core.InitLogger()
	// 初始化数据库
	global.DB = core.InitGorm()
	// 初始化路由
	router := routers.InitRouter()

	global.Log.Info("初始化完成")
	router.Run(global.Conf.Sys.Addr())
}

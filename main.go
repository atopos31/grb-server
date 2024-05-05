package main

import (
	"flag"
	"fmt"
	"gvb/core"
	"gvb/global"
)

func main() {
	configDir := flag.String("config", "./config/dev.yaml", "config dir")
	flag.Parse()
	// 初始化配置
	core.InitConf(*configDir)
	fmt.Println(global.Conf)
	// 初始化数据库
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}

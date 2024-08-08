package app

import (
	"gvb/config"

	"github.com/sirupsen/logrus"
)

var (
	// Conf 全局配置
	Conf *config.Config
	// Log 日志
	Log *logrus.Logger
)

func PrintLogo() {
	println(`
	
   _____   _____    ____              _____                                     
  / ____| |  __ \  |  _ \            / ____|                                    
 | |  __  | |__) | | |_) |  ______  | (___     ___   _ __  __   __   ___   _ __ 
 | | |_ | |  _  /  |  _ <  |______|  \___ \   / _ \ | '__| \ \ / /  / _ \ | '__|
 | |__| | | | \ \  | |_) |           ____) | |  __/ | |     \ V /  |  __/ | |   
  \_____| |_|  \_\ |____/           |_____/   \___| |_|      \_/    \___| |_|   
                                                                                
                                                                                

	`)
}

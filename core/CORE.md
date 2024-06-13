整个项目依赖的核心，无法初始化则程序无法启动
```golang
// 初始化配置信息
global.Conf = core.InitConf(*configDir)
// 初始化日志
global.Log = core.InitLogger(global.Conf.Logger)
// 初始化数据库
db := core.InitGorm(global.Conf.Mysql)
// 初始化redis
cache := core.InitRedis(global.Conf.Redis)
// 初始化站点信息
siteInfo := core.InitSiteInfo(global.Conf.Sys.SiteInfoPath)
```
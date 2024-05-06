package routers

import (
	"gvb/api/user"
	"gvb/config"
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(config config.System) *gin.Engine {
	gin.SetMode(config.Env)
	router := gin.Default()
	// 中间件
	router.Use(middleware.Cors(config.Origin))
	// 注册用户相关路由
	user.RegisRouter(router)

	return router
}

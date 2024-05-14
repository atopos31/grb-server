package routers

import (
	"gvb/api/article"
	"gvb/api/user"
	"gvb/config"
	"gvb/middleware"

	_ "gvb/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(config config.System) *gin.Engine {
	gin.SetMode(config.Env)
	router := gin.Default()
	// 中间件
	router.Use(middleware.Cors(config.Origin))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routerGroup := router.Group("/api")
	// 注册用户相关路由
	user.RegisRouter(routerGroup)
	article.RegisRouter(routerGroup)

	return router
}

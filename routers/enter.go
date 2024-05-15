package routers

import (
	"gvb/api/article"
	"gvb/api/user"
	"gvb/config"
	"gvb/middleware"

	docs "gvb/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(config config.System) *gin.Engine {
	gin.SetMode(config.Env)
	router := gin.Default()
	// swagger文档
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 中间件
	router.Use(middleware.Cors(config.Origin))

	routerGroup := router.Group("/api")
	// 注册用户相关路由
	user.RegisRouter(routerGroup)
	article.RegisRouter(routerGroup)

	return router
}

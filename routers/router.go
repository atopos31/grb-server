package routers

import (
	"gvb/api/article"
	"gvb/api/category"
	"gvb/api/comment"
	"gvb/api/info"
	"gvb/api/oss"
	"gvb/api/system"
	"gvb/api/tag"
	"gvb/api/user"
	"gvb/config"
	"gvb/middleware"

	docs "gvb/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New(config config.System) *gin.Engine {
	gin.SetMode(config.Env)
	router := gin.Default()
	// swagger文档
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 中间件
	router.Use(middleware.Cors(config.Origin))
	// base URL
	routerGroup := router.Group("/api")

	user.RegisRouter(routerGroup)
	article.RegisRouter(routerGroup)
	category.RegisRouter(routerGroup)
	comment.RegisRouter(routerGroup)
	tag.RegisRouter(routerGroup)
	oss.RegisRouter(routerGroup)
	info.RegisRouter(routerGroup)
	system.RegisRouter(routerGroup)

	return router
}

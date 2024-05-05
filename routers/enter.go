package routers

import (
	"gvb/api/user"
	"gvb/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Conf.Sys.Env)
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 注册用户相关路由
	user.RegisRouter(router)

	return router
}

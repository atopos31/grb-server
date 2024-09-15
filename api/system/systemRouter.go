package system

import (
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisRouter(c *gin.RouterGroup) {
	hostApi := c.Group("/system")
	hostApi.Use(middleware.Auth())
	{
		hostApi.GET("/info", getInfo)
		hostApi.GET("/cmn", getCmn)
	}
}

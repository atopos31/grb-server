package host

import (
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisRouter(c *gin.RouterGroup) {
	hostApi := c.Group("/host")
	hostApi.Use(middleware.Auth())
	{
		hostApi.GET("/get", GetHostInfo)
	}
}

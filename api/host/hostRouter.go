package host

import "github.com/gin-gonic/gin"

func RegisRouter(c *gin.RouterGroup) {
	hostApi := c.Group("/host")
	{
		hostApi.GET("/get", GetHostInfo)
	}
}

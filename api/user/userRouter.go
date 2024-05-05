package user

import "github.com/gin-gonic/gin"

func RegisRouter(c *gin.Engine) {
	userApi := c.Group("/user")
	userApi.GET("/login", Login)
}

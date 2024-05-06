package user

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.Engine) {
	userApi := r.Group("/user")
	{
		userApi.GET("/login", Login)
	}
}

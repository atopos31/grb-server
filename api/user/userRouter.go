package user

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	userApi := r.Group("/user")
	{
		userApi.POST("/login", login)
	}
}

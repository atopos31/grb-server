package user

import (
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisRouter(r *gin.RouterGroup) {
	userApi := r.Group("/user")
	{
		userApi.POST("/login", login)
		
		userManageApi := userApi.Group("/manage")
		userManageApi.Use(middleware.Auth())
		{
			userManageApi.GET("/info", getUserInfo)
		}
	}
}

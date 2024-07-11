package oss

import (
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisRouter(r *gin.RouterGroup) {
	ossApi := r.Group("/oss")
	ossApi.Use(middleware.Auth())
	{
		ossApi.GET("/uptoken", uptoken)
	}
}

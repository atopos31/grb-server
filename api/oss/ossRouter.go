package oss

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	ossApi := r.Group("/oss")
	{
		ossApi.GET("/uptoken", uptoken)
	}
}

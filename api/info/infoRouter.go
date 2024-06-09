package info

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	infoApi := r.Group("/info")
	{
		infoApi.GET("/site", GetSiteInfo)
		infoApi.GET("/basic", GetBasicInfo)
	}
}

package info

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	infoApi := r.Group("/info")
	{
		infoApi.GET("/siteinfo", GetSiteInfo)
	}
}

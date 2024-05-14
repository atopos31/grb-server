package article

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	articleApi := r.Group("/article")
	{
		articleApi.GET("/list")
		articleApi.POST("/create", create)
	}
}

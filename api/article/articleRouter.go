package article

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.Engine) {
	articleApi := r.Group("/article")
	{
		articleApi.GET("/list", )
	}
}
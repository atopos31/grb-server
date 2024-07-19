package comment

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	commentApi := r.Group("/comment")
	{
		commentApi.POST("/create", create)
		commentApi.GET("/list/:uuid", getList)
	}
}

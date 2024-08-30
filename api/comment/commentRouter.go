package comment

import (
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisRouter(r *gin.RouterGroup) {
	commentApi := r.Group("/comment")
	{
		commentApi.POST("/create", create)
		commentApi.GET("/list/:uuid", getArticleCommentsList)

		commentManageApi := commentApi.Group("/manage")
		commentManageApi.Use(middleware.Auth())
		{
			commentManageApi.GET("/list/:status", getList)
			commentManageApi.PUT("/status/:id", update)
		}
	}
}

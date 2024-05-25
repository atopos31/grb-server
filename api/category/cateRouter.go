package category

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	cateApi := r.Group("/category")
	{
		cateApi.POST("/create", create)
		cateApi.GET("/list", getList)
		cateApi.PUT("/update", update)
		cateApi.DELETE("/delete/:id", delete)
	}
}

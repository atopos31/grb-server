package tag

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	tagApi := r.Group("/tag")
	{
		tagApi.GET("/list", getList)
		tagApi.GET("/hotlist/:size", getHotList)
		tagApi.POST("/create", create)
		tagApi.PUT("/update", update)
		tagApi.DELETE("/delete/:id", delete)
	}

}

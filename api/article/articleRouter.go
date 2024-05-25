package article

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	articleApi := r.Group("/article")
	{
		// 获取文章列表
		articleApi.GET("/list", getList)
		// 创建文章
		articleApi.POST("/create", create)
		// 获取文章详情
		articleApi.GET("/get", getByUuid)
		// 更新文章
		articleApi.PUT("/update/:uuid", update)
		// 删除文章
		articleApi.DELETE("/delete/:uuid", delete)
	}
}

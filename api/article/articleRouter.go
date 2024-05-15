package article

import "github.com/gin-gonic/gin"

func RegisRouter(r *gin.RouterGroup) {
	articleApi := r.Group("/article")
	{
		// 文章列表
		articleApi.GET("/list", getList)
		// 创建文章
		articleApi.POST("/create", create)
		// 获取文章详情
		articleApi.GET("/get", getByUuid)
	}
}

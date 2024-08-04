package article

import (
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisRouter(r *gin.RouterGroup) {
	articleApi := r.Group("/article")
	{
		// 获取文章列表
		articleApi.GET("/list", getList)
		// 获取文章详情
		articleApi.GET("/get", getByUuid)
		// 基于搜索引擎搜索
		articleApi.GET("/search", searchByQuery)

		//后台接口
		articleManageApi := articleApi.Group("/manage")
		articleManageApi.Use(middleware.Auth())
		{
			// 创建文章
			articleManageApi.POST("/create", create)
			// 全量更新文章
			articleManageApi.PUT("/update/:uuid", update)
			// 局部更新文章
			articleManageApi.PATCH("/update/:uuid", updatesection)
			// 删除文章
			articleManageApi.DELETE("/delete/:uuid", delete)
			// 获取管理文章列表
			articleManageApi.GET("/list", getManageList)
		}

	}
}

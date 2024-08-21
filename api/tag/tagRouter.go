package tag

import (
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisRouter(r *gin.RouterGroup) {
	tagApi := r.Group("/tag")
	{
		tagApi.GET("/list", getList)
		// 获取热门标签
		tagApi.GET("/hotlist/:size", getHotList)
		// 管理
		tagManageApi := tagApi.Group("/manage")
		tagManageApi.Use(middleware.Auth())
		{
			tagManageApi.GET("/list", getManageList)
			tagManageApi.POST("/create", create)
			tagManageApi.PUT("/update", update)
			tagManageApi.DELETE("/delete/:id", delete)
		}
	}

}

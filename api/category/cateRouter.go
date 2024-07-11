package category

import (
	"gvb/middleware"

	"github.com/gin-gonic/gin"
)

func RegisRouter(r *gin.RouterGroup) {
	cateApi := r.Group("/category")
	{
		cateApi.GET("/list", getList)
		cateManageApi := cateApi.Group("/manage")
		cateManageApi.Use(middleware.Auth())
		{
			cateManageApi.GET("/list", getManageList)
			cateManageApi.POST("/create", create)
			cateManageApi.PUT("/update", update)
			cateManageApi.DELETE("/delete/:id", delete)
		}
	}
}

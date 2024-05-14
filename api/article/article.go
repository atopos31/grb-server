package article

import (
	"gvb/models/errcode"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/service"

	"github.com/gin-gonic/gin"
)

var response = res.NewResponse()

// @Summary 创建文章
// @Description 创建文章
// @Tags 文章
// @Accept json
// @Produce json
// @Param data body req.Article true "文章信息"
// @Success 200 {object} res.Response
// @Router /api/article/create [post]
func create(c *gin.Context) {
	articleReq := req.Article{}

	if err := c.ShouldBindJSON(&articleReq); err != nil {
		response.ErrorRaw(c, err)
		return
	}

	if err := service.Svc.ArticleService.Create(&articleReq); err != nil {
		response.Error(c, errcode.ErrInternalServer)
		return
	}
	response.Success(c, nil)
}

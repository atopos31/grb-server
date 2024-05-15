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
// @Router /article/create [post]
func create(c *gin.Context) {
	articleReq := new(req.Article)

	if err := c.ShouldBindJSON(articleReq); err != nil {
		response.ErrorRaw(c, err)
		return
	}

	if err := service.Svc.ArticleService.Create(articleReq); err != nil {
		response.Error(c, errcode.ErrInternalServer)
		return
	}
	response.Success(c, nil)
}

// @Summary 获取文章列表
// @Description 获取文章列表
// @Tags 文章
// @Accept json
// @Produce json
// @Param data query req.ArticleList true "文章列表"
// @Success 200 {object} res.Response
// @Router /article/list [get]
func getList(c *gin.Context) {
	articleList := new(req.ArticleList)
	if err := c.ShouldBindQuery(articleList); err != nil {
		response.ErrorRaw(c, err)
		return
	}
	list, err := service.Svc.ArticleService.GetList(articleList)
	if err != nil {
		response.Error(c, errcode.ErrInternalServer)
		return
	}
	response.Success(c, list)
}

// @Summary 获取文章详情
// @Description 获取文章详情
// @Tags 文章
// @Accept json
// @Produce json
// @Param uuid query string true "文章uuid"
// @Success 200 {object} res.Response
// @Router /article/get [get]
func getByUuid(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		response.Error(c, errcode.ErrInvalidParam)
		return
	}
	article, err := service.Svc.ArticleService.GetByUuid(uuid)
	if err != nil {
		response.Error(c, errcode.ErrInternalServer)
		return
	}
	response.Success(c, article)
}

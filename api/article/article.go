package article

import (
	"gvb/models/errcode"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/service"

	"github.com/gin-gonic/gin"
)

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
		res.ErrorRaw(c, err)
		return
	}

	if err := service.Svc.ArticleService.Create(articleReq); err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, nil)
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
		res.ErrorRaw(c, err)
		return
	}
	list, err := service.Svc.ArticleService.GetList(articleList)
	if err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, list)
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
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	article, err := service.Svc.ArticleService.GetByUuid(uuid)
	if err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, article)
}

// @Summary 删除文章
// @Description 删除文章
// @Tags 文章
// @Accept multipart/form-data
// @Produce json
// @Param uuid formData string true "文章uuid"
// @Success 200 {object} res.Response
// @Router /article/delete [post]
func delete(c *gin.Context) {
	uuid := c.PostForm("uuid")
	if uuid == "" {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	if err := service.Svc.ArticleService.DeleteByUuid(uuid); err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, nil)
}

// @Summary 更新文章
// @Tags 文章
// @Accept json
// @Produce json
// @Param article body req.Article true "文章信息"
// @Success 200 {object} res.Response
// @Router /article/update [post]
func update(c *gin.Context) {
	articleReq := new(req.Article)

	if err := c.ShouldBindJSON(articleReq); err != nil {
		res.ErrorRaw(c, err)
		return
	}
	if articleReq.Uuid == 0 {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	if err := service.Svc.ArticleService.Update(articleReq); err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, nil)
}

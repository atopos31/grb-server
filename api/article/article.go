package article

import (
	"gvb/global"
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
// @Router /article/manage/create [post]
func create(c *gin.Context) {
	articleReq := new(req.Article)

	if err := c.ShouldBindJSON(articleReq); err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}

	resCreate, err := service.Svc.ArticleService.Create(articleReq)
	if err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}

	res.Success(c, resCreate)
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
		res.Error(c, err)
		return
	}
	list, err := service.Svc.ArticleService.GetList(articleList)
	if err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, list)
}

// @Summary 获取文章列表(管理员)
// @Description 获取文章列表(管理员)
// @Tags 文章
// @Accept json
// @Produce json
// @Param data query req.ArticleList true "文章列表"
// @Success 200 {object} res.Response
// @Router /article/manage/list [get]
func getManageList(c *gin.Context) {
	articleList := new(req.ArticleList)
	if err := c.ShouldBindQuery(articleList); err != nil {
		res.Error(c, err)
		return
	}
	list, err := service.Svc.ArticleService.GetManageList(articleList)
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
// @Param uuid path string true "文章uuid"
// @Success 200 {object} res.Response
// @Router /article/manage/delete/:uuid [delete]
func delete(c *gin.Context) {
	uuid := c.Param("uuid")
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
// @Param uuid path string true "文章uuid"
// @Param article body req.Article true "文章信息"
// @Success 200 {object} res.Response
// @Router /article/manage/update/:uuid [put]
func update(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}

	articleReq := new(req.Article)
	if err := c.ShouldBindJSON(articleReq); err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	global.Log.Info(articleReq)
	resUpdate, err := service.Svc.ArticleService.Update(articleReq, uuid)
	if err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, resUpdate)
}

// @Summary 局部更新文章
// @Tags 文章
// @Accept json
// @Produce json
// @Param uuid path string true "文章uuid"
// @Param article body req.ArticleSertion true "文章信息"
// @Success 200 {object} res.Response
// @Router /article/manage/update/:uuid [patch]
func updatesection(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	articleSertion := new(req.ArticleSertion)
	if err := c.ShouldBindJSON(articleSertion); err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	if err := service.Svc.ArticleService.UpdateSectionByUuid(uuid, articleSertion); err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, nil)
}

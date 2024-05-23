package category

import (
	"gvb/models/errcode"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 创建分类
// @Tags 分类
// @Produce json
// @Param data body req.Cate true "创建分类"
// @Success 200 {object} res.Response
// @Router /category/create [post]
func create(c *gin.Context) {
	cateReq := new(req.Cate)

	if err := c.ShouldBindJSON(cateReq); err != nil {
		res.ErrorRaw(c, err)
		return
	}
	if err := service.Svc.CateService.Create(cateReq); err != nil {
		res.ErrorRaw(c, err)
		return
	}

	res.Success(c, nil)
}

// @Summary 获取分类列表
// @Tags 分类
// @Produce json
// @Success 200 {object} res.Response
// @Router /category/list [get]
func getList(c *gin.Context) {
	cateList, err := service.Svc.CateService.GetList()
	if err != nil {
		res.ErrorRaw(c, err)
		return
	}
	res.Success(c, cateList)
}

// @Summary 更新分类
// @Tags 分类
// @Produce json
// @Param data body req.Cate true "更新分类"
// @Success 200 {object} res.Response
// @Router /category/update [post]
func update(c *gin.Context) {
	cateReq := new(req.Cate)
	if err := c.ShouldBindJSON(cateReq); err != nil {
		res.ErrorRaw(c, err)
		return
	}
	if err := service.Svc.CateService.Update(cateReq); err != nil {
		res.ErrorRaw(c, err)
		return
	}
	res.Success(c, nil)
}

// @Summary 删除分类
// @Tags 分类
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} res.Response
// @Router /category/delete/{id} [post]
func delete(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	if err := service.Svc.CateService.Delete(uint(idint)); err != nil {
		res.ErrorRaw(c, err)
		return
	}
	res.Success(c, nil)
}

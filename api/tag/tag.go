package tag

import (
	"gvb/models/errcode"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Tags 标签
// @Summary 创建标签
// @Description 创建标签
// @Produce json
// @Param data body req.Tag true "标签"
// @Success 200 {object} res.Response
// @Router /tag/create [post]
func create(c *gin.Context) {
	tag := new(req.Tag)
	if err := c.ShouldBindJSON(tag); err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	id, err := service.Svc.TagService.Create(tag)
	if err != nil {
		res.Error(c, err)
		return
	}
	res.Success(c, id)
}

// @Summary 获取标签列表
// @Description 获取标签列表
// @Tags 标签
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Router /tag/list [get]
func getList(c *gin.Context) {
	tags, err := service.Svc.TagService.GetList()
	if err != nil {
		res.Error(c, err)
		return
	}
	res.Success(c, tags)
}

// @Summary 更新标签
// @Description 更新标签
// @Tags 标签
// @Accept json
// @Produce json
// @Param tag body req.Tag true "标签"
// @Success 200 {object} res.Response
// @Router /tag/update [put]
func update(c *gin.Context) {
	tag := new(req.Tag)
	if err := c.ShouldBindJSON(tag); err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	if err := service.Svc.TagService.Update(tag); err != nil {
		res.Error(c, err)
		return
	}
	res.Success(c, nil)
}

// @Summary 删除标签
// @Tags 标签
// @Produce json
// @Param id formData int true "标签ID"
// @Success 200 {object} res.Response
// @Router /tag/delete/:id [delete]
func delete(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	if err := service.Svc.TagService.Delete(uint(idint)); err != nil {
		res.Error(c, err)
		return
	}
	res.Success(c, nil)
}

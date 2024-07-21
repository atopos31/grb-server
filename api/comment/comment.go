package comment

import (
	"gvb/models/errcode"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/service"
	"gvb/utils/conver"

	"github.com/gin-gonic/gin"
)

// @Summary 创建评论
// @Tags 评论
// @Description 创建评论
// @Accept json
// @Produce json
// @Param data body req.Comment true "评论"
// @Success 200 {object} res.Response{data=res.Response}
// @Router /comment/create [post]
func create(c *gin.Context) {
	commentReq := new(req.Comment)
	if err := c.ShouldBindJSON(commentReq); err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	if err := service.Svc.CommentService.Create(commentReq); err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, nil)
}

// @Summary 获取评论列表
// @Tags 评论
// @Description 获取评论列表
// @Accept json
// @Produce json
// @Param data body req.Comment true "评论"
// @Success 200 {object} res.Response{data=[]res.Comment}
// @Router /comment/list/:uuid [get]
func getList(c *gin.Context) {
	uuidStr := c.Param("uuid")
	if uuidStr == "" {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}

	uuid, err := conver.StrToUint32(uuidStr)
	if err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}

	list, err := service.Svc.CommentService.GetList(uuid)
	if err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, list)
}

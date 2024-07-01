package user

import (
	"gvb/models/errcode"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/service"

	"github.com/gin-gonic/gin"
)

// @Summary 登录
// @Tags 用户
// @Produce json
// @Param data body req.User true "用户名密码"
// @Success 200 {object} res.Response
// @Router /user/login [post]
func login(c *gin.Context) {
	user := new(req.User)
	if err := c.ShouldBindJSON(user); err != nil {
		res.Error(c, errcode.ErrInvalidParam)
		return
	}
	token, err := service.Svc.UserService.Login(user)
	if err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}
	res.Success(c, token)
}

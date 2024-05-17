package user

import (
	"gvb/models/errcode"
	"gvb/models/res"
	"gvb/service"

	"github.com/gin-gonic/gin"
)

// 登录 handler
func login(c *gin.Context) {
	if err := service.Svc.UserService.Login(); err == nil {
		res.Success(c, nil)
		return
	}
	res.Error(c, errcode.ErrInternalServer)
}

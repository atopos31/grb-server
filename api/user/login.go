package user

import (
	"gvb/models/errcode"
	"gvb/service"

	"github.com/gin-gonic/gin"
)

// 登录 handler
func Login(c *gin.Context) {
	if err := service.Svc.UserService.Login(); err == nil {
		response.Success(c, nil)
		return
	}
	response.Error(c, errcode.ErrInternalServer)
}

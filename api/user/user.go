package user

import (
	"gvb/models/errcode"
	"gvb/models/res"
	"gvb/service"

	"github.com/gin-gonic/gin"
)

var response = res.NewResponse()

// 登录 handler
func login(c *gin.Context) {
	if err := service.Svc.UserService.Login(); err == nil {
		response.Success(c, nil)
		return
	}
	response.Error(c, errcode.ErrInternalServer)
}

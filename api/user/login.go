package user

import (
	"github.com/gin-gonic/gin"
)

// 登录 handler
func Login(c *gin.Context) {
	response.Success(c, "登录成功")
}

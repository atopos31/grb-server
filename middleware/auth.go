package middleware

import (
	"fmt"
	"gvb/global"
	"gvb/models/errcode"
	"gvb/models/res"
	"gvb/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		header := c.Request.Header.Get("Authorization")
		if len(header) == 0 {
			res.Error(c, errcode.ErrAccessDenied)
			c.Abort()
			return
		}

		var token string
		if _, err := fmt.Sscanf(header, "Bearer %s", &token); err != nil {
			res.Error(c, errcode.ErrAccessDenied)
			c.Abort()
			return
		}

		payload, err := jwt.Parse(token, global.Conf.Jwt.Secret)
		if err != nil {
			res.Error(c, errcode.ErrAccessDenied)
			c.Abort()
			return
		}

		c.Set("username", payload.Username)
		return
	}
}

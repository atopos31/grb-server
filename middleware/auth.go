package middleware

import (
	"fmt"
	"gvb/models/errcode"
	"gvb/models/res"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if len(header) == 0 {
			res.Error(c, errcode.ErrAccessDenied)
			c.Abort()
			return
		}

		var token string
		_, err := fmt.Sscanf(header, "Bearer %s", &token)
		if err != nil {
			res.Error(c, errcode.ErrAccessDenied)
			c.Abort()
			return
		}

		// TODO: 验证token
	}
}

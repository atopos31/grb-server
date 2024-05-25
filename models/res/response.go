package res

import (
	"net/http"

	"gvb/models/errcode"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, Response{
		Code:    errcode.Success.Code(),
		Message: errcode.Success.Msg(),
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	// 断言
	codeErr, ok := err.(errcode.Error)
	if ok {
		c.JSON(http.StatusOK, Response{
			Code:    codeErr.Code(),
			Message: codeErr.Msg(),
			Data:    gin.H{},
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Code:    9999,
			Message: err.Error(),
			Data:    gin.H{},
		})
	}
}

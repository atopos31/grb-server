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

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Success(c *gin.Context, data any) {
	if data == nil {
		data = gin.H{}
	}
	c.JSON(http.StatusOK, Response{
		Code:    errcode.Success.Code(),
		Message: errcode.Success.Msg(),
		Data:    data,
	})
}

func (r *Response) Error(c *gin.Context, err *errcode.Error) {
	c.JSON(http.StatusOK, Response{
		Code:    err.Code(),
		Message: err.Msg(),
		Data:    gin.H{},
	})
}

func (r *Response) ErrorRaw(c *gin.Context, err error) {
	c.JSON(http.StatusOK, Response{
		Code:    9999,
		Message: err.Error(),
		Data:    gin.H{},
	})
}

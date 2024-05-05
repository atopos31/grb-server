package res

import (
	"gvb/models/errcode"
	"net/http"

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

func (r *Response) Error(c *gin.Context, err errcode.Error) {
	c.JSON(http.StatusOK, Response{
		Code:    err.Code(),
		Message: err.Msg(),
		Data:    gin.H{},
	})
}

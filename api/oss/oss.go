package oss

import (
	"gvb/models/errcode"
	"gvb/models/res"
	"gvb/service"

	"github.com/gin-gonic/gin"
)

// @Summary OSS
// @Description 获取oss上传token
// @Tags OSS
// @Produce json
// @Success 200 {object} res.OssConfig
// @Router /oss/uptoken [get]
func uptoken(c *gin.Context) {
	resUpToken, err := service.Svc.OssService.GetUploadToken()
	if err != nil {
		res.Error(c, errcode.ErrInternalServer)
		return
	}

	res.Success(c, resUpToken)
}

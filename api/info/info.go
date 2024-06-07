package info

import (
	"gvb/models/res"
	"gvb/service"

	"github.com/gin-gonic/gin"
)

// @Tags 站点信息
// @Summary 获取站点信息
// @Produce json
// @Success 200 {object} res.SiteInfo
// @Router /info/siteinfo [get]
func GetSiteInfo(c *gin.Context) {
	info, err := service.Svc.SiteInfoService.GetSiteInfo(c)
	if err != nil {
		res.Error(c, err)
		return
	}
	res.Success(c, info)
}

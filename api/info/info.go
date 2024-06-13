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
// @Router /info/site [get]
func GetSiteInfo(c *gin.Context) {
	info, err := service.Svc.SiteInfoService.GetSiteInfo(c)
	if err != nil {
		res.Error(c, err)
		return
	}
	res.Success(c, info)
}

// @Tags 站点信息
// @Summary 获取站点基本信息
// @Produce json
// @Success 200 {object} res.BasicInfo
// @Router /info/basic [get]
func GetBasicInfo(c *gin.Context) {
	res.Success(c, service.Svc.SiteInfoService.GetBasicInfo())
}

// @Tags 站点信息
// @Summary 获取站点徽标信息
// @Produce json
// @Success 200 {object} []site.Badge
// @Router /info/badges [get]
func GetBadges(c *gin.Context) {
	res.Success(c, service.Svc.SiteInfoService.GetBadges())
}

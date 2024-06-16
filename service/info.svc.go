package service

import (
	"gvb/models/entity"
	"gvb/models/res"
	"gvb/site"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

const viewkey = "view:"

type SiteInfoService struct {
	db           *gorm.DB
	cache        *redis.Client
	siteInfoPath string
	siteInfo     *site.SieInfo
}

func NewSiteInfoService(db *gorm.DB, cache *redis.Client, siteInfoPath string) *SiteInfoService {
	siteInfo, err := loadInfo(siteInfoPath)
	if err != nil {
		panic(err)
	}
	return &SiteInfoService{
		db:       db,
		cache:    cache,
		siteInfo: siteInfo,
	}
}

func (s *SiteInfoService) GetSiteInfo(ctx *gin.Context) (*res.SiteInfo, error) {
	siteInfo := new(res.SiteInfo)
	if err := s.db.Model(&entity.Article{}).
		Where("status = ?", 1).
		Count(&siteInfo.ArticleCount).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&entity.Category{}).Count(&siteInfo.CategoryCount).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&entity.Tag{}).Count(&siteInfo.TagCount).Error; err != nil {
		return nil, err
	}

	exists, err := s.cache.Exists(ctx, viewkey+ctx.ClientIP()).Result()
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		siteInfo.ViewsCount, err = s.cache.Incr(ctx, viewkey).Result()
		if err != nil {
			return nil, err
		}
		if err := s.cache.Set(ctx, viewkey+ctx.ClientIP(), 1, 24*time.Hour).Err(); err != nil {
			return nil, err
		}
	} else {
		siteInfo.ViewsCount, err = s.cache.Get(ctx, viewkey).Int64()
		if err != nil {
			return nil, err
		}
	}
	siteInfo.Record = s.siteInfo.Site.Record
	return siteInfo, nil
}

func (s *SiteInfoService) GetBasicInfo() res.BasicInfo {
	return res.BasicInfo{
		Social: s.siteInfo.Social,
		User:   s.siteInfo.User,
	}
}

func (s *SiteInfoService) GetBadges() []site.Badge {
	return s.siteInfo.Badges
}

func loadInfo(filePath string) (*site.SieInfo, error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	siteInfo := new(site.SieInfo)
	if err := yaml.Unmarshal(file, siteInfo); err != nil {
		return nil, err
	}

	return siteInfo, nil
}

func reloadInfo(filePath string, siteInfo site.SieInfo) (*site.SieInfo, error) {
	data, err := yaml.Marshal(siteInfo)
	if err != nil {
		return nil, err
	}

	if err = ioutil.WriteFile(filePath, data, 0644); err != nil {
		return nil, err
	}
	return &siteInfo, nil
}

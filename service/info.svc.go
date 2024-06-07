package service

import (
	"gvb/models/entity"
	"gvb/models/res"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const viewkey = "view:"

type SiteInfoService struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewSiteInfoService(db *gorm.DB, cache *redis.Client) *SiteInfoService {
	return &SiteInfoService{
		db:    db,
		cache: cache,
	}
}

func (s *SiteInfoService) GetSiteInfo(ctx *gin.Context) (*res.SiteInfo, error) {
	siteInfo := new(res.SiteInfo)
	if err := s.db.Model(&entity.Article{}).Count(&siteInfo.ArticleCount).Error; err != nil {
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
		siteInfo.ViewsCount, err = s.cache.Incr(ctx, viewkey+ctx.ClientIP()).Result()
		if err != nil {
			return nil, err
		}
		if err := s.cache.Expire(ctx, viewkey+ctx.ClientIP(), 24*60*60).Err(); err != nil {
			return nil, err
		}
	} else {
		siteInfo.ViewsCount, err = s.cache.Get(ctx, viewkey+ctx.ClientIP()).Int64()
		if err != nil {
			return nil, err
		}
	}
	siteInfo.Record = "辽ICP备2022010174号"
	return siteInfo, nil
}

package service

import (
	"gvb/config"
	"gvb/dao"
	"gvb/service/inter"
	"gvb/site"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var Svc *Service

type Service struct {
	UserService     *UserService
	ArticleService  *ArticleService
	CateService     *CateService
	TagService      *TagService
	SiteInfoService *SiteInfoService
	OssService      inter.OssService // oss接口
	AiService       inter.AiService  // ai接口
}

func New(db *gorm.DB, cache *redis.Client, ossConfig config.Oss, aiConfig config.Ai, siteInfo site.SieInfo) *Service {
	userRepo := dao.NewUserRepo(db)
	articleRepo := dao.NewArticleRepo(db, cache)
	tagRepo := dao.NewTagRepo(db)
	cateRepo := dao.NewCateRepo(db)

	//预留接口 实现可拓展 可选择不同Oss服务注入
	ossSvc := NewOssQinui(ossConfig.OssQiniu)
	aiSvc := NewAiHunyuan(aiConfig.Hunyuan)

	return &Service{
		UserService:     NewUserService(userRepo),
		ArticleService:  NewArticleService(articleRepo, tagRepo),
		CateService:     NewCateService(cateRepo),
		TagService:      NewTagService(tagRepo),
		SiteInfoService: NewSiteInfoService(db, cache, siteInfo),
		OssService:      ossSvc,
		AiService:       aiSvc,
	}
}

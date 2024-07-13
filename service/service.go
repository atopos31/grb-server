package service

import (
	"gvb/config"
	"gvb/dao"
	"gvb/service/inter"

	"github.com/meilisearch/meilisearch-go"
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
	OssService      inter.OssService // OSS接口
	AiService       inter.AiService  // AI接口
}

func New(db *gorm.DB, search *meilisearch.Client, cache *redis.Client, ossConfig config.Oss, aiConfig config.Ai, siteInfoPath string) *Service {
	userRepo := dao.NewUserRepo(db)
	articleRepo := dao.NewArticleRepo(db, cache, search)
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
		SiteInfoService: NewSiteInfoService(db, cache, siteInfoPath),
		OssService:      ossSvc,
		AiService:       aiSvc,
	}
}

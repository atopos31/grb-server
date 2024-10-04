package service

import (
	"gvb/app"
	"gvb/config"
	"gvb/core"
	"gvb/dao"
	"gvb/service/inter"
)

var Svc *Service

type Service struct {
	UserService     *UserService
	ArticleService  *ArticleService
	CateService     *CateService
	TagService      *TagService
	CommentService  *CommentService
	SiteInfoService *SiteInfoService
	OssService      inter.OssService // OSS接口
	AiService       inter.AiService  // AI接口
}

const hunyuanAI = "hunyuan"
const qianfanAI = "qianfan"

func New(config config.Config) *Service {
	db := core.NewGormDB(config.Mysql)
	search := core.NewMeiliSearchClient(config.Meilisearch)
	cache := core.NewRedisCache(config.Redis)

	userRepo := dao.NewUserRepo(db)
	articleRepo := dao.NewArticleRepo(db, cache, search,config.Meilisearch.Index)
	tagRepo := dao.NewTagRepo(db)
	cateRepo := dao.NewCateRepo(db)
	commentRepo := dao.NewCommentRepo(db)

	//预留接口 实现可拓展 可选择不同Oss服务注入
	ossSvc := NewOssQinui(config.Oss.OssQiniu)
	aiSvc := NewAiSvcByConfig(config.Ai)

	siteInfoPath := config.Sys.SiteInfoPath

	return &Service{
		UserService:     NewUserService(userRepo),
		ArticleService:  NewArticleService(articleRepo, tagRepo),
		CateService:     NewCateService(cateRepo),
		TagService:      NewTagService(tagRepo),
		CommentService:  NewCommentService(commentRepo),
		SiteInfoService: NewSiteInfoService(db, cache, siteInfoPath),
		OssService:      ossSvc,
		AiService:       aiSvc,
	}
}

func NewAiSvcByConfig(config config.Ai) inter.AiService {
	switch config.Use {
	case hunyuanAI:
		return NewAiHunyuan(config.Hunyuan)
	case qianfanAI:
		return NewAiQianfan(config.Qianfan)
	default:
		app.Log.Error("AI service create failed", config)
		return NewAiDefault(config)
	}
}

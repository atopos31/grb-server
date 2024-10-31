package service

import (
	"gvb/app"
	"gvb/config"
	"gvb/service/inter"

	"github.com/google/wire"
)

var Svc *Service

var ProvideSet = wire.NewSet(
	NewService,
	NewOssQinui,
	NewUserService,
	NewArticleService,
	NewCateService,
	NewTagService,
	NewCommentService,
	NewSiteInfoService,
	NewAiSvcByConfig,
)

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

func NewService(config config.Config, userSvc *UserService, articleSvc *ArticleService, cateSvc *CateService, tagSvc *TagService, commentSvc *CommentService, siteSvc *SiteInfoService, ossSvc inter.OssService, aiSvc inter.AiService) *Service {
	return &Service{
		UserService:     userSvc,
		ArticleService:  articleSvc,
		CateService:     cateSvc,
		TagService:      tagSvc,
		CommentService:  commentSvc,
		SiteInfoService: siteSvc,
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
		app.Log.Errorf("AI service create failed %v", config)
		return NewAiDefault(config)
	}
}

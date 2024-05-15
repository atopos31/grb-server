package service

import (
	"gvb/dao"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var Svc *Service

type Service struct {
	UserService    *UserService
	ArticleService *ArticleService
}

func New(db *gorm.DB, cache *redis.Client) *Service {
	userRepo := dao.NewUserRepo(db)
	articleRepo := dao.NewArticleRepo(db, cache)
	tagRepo := dao.NewTagRepo(db)

	return &Service{
		UserService:    NewUserService(userRepo),
		ArticleService: NewArticleService(articleRepo, tagRepo),
	}
}

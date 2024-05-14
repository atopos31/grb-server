package dao

import (
	"gvb/models/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ArticleRepo struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewArticleRepo(db *gorm.DB, cache *redis.Client) *ArticleRepo {
	return &ArticleRepo{db: db, cache: cache}
}

func (a *ArticleRepo) Create(article entity.Article) error {
	return a.db.Create(&article).Error
}

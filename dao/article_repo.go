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

func (a *ArticleRepo) GetList(pageSize, pageNum int) ([]entity.Article, error) {
	var articles []entity.Article
	err := a.db.Debug().Preload("Tags").Preload("Category").
		Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("Created_At DESC").Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, err
}

func (a *ArticleRepo) GetByUuid(uuid string) (entity.Article, error) {
	var article entity.Article
	err := a.db.Debug().Preload("Tags").Preload("Category").Where("uuid = ?", uuid).First(&article).Error
	if err != nil {
		return article, err
	}
	return article, err
}

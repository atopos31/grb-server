package dao

import (
	"gvb/models/entity"
	"gvb/models/req"

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

// GetList 获取文章列表
func (a *ArticleRepo) GetList(pageSize, pageNum int) ([]entity.Article, error) {
	var articles []entity.Article
	err := a.db.Preload("Tags").Preload("Category").
		Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("Created_At DESC").Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, err
}

// GetByUuid 根据uuid获取文章
func (a *ArticleRepo) GetByUuid(uuid string) (entity.Article, error) {
	var article entity.Article
	err := a.db.Preload("Tags").Preload("Category").Where("uuid = ?", uuid).First(&article).Error
	if err != nil {
		return article, err
	}
	return article, err
}

// 根据uuid删除文章
func (a *ArticleRepo) DeleteByUuid(uuid string) error {
	var article entity.Article
	if err := a.db.Preload("Tags").Where("uuid = ?", uuid).First(&article).Error; err != nil {
		return err
	}
	// 删除文章对应的标签
	if err := a.db.Model(&article).Association("Tags").Clear(); err != nil {
		return err
	}

	return a.db.Delete(&article).Error
}

// 根据uuid更新文章
func (a *ArticleRepo) UpdateByUuid(newArticle *req.Article, uuid string) (*entity.Article, error) {
	tx := a.db.Begin()
	article := new(entity.Article)
	if err := tx.Where("uuid = ?", uuid).First(&article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	article.Title = newArticle.Title
	article.Content = newArticle.Content
	article.CoverImage = newArticle.CoverImage
	article.CategoryID = newArticle.CategoryID
	article.Top = newArticle.Top
	article.Status = newArticle.Status

	if err := tx.Save(article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return article, tx.Commit().Error
}

// 更新文章的标签
func (a *ArticleRepo) UpdateTags(article *entity.Article, tags []entity.Tag) error {
	return a.db.Model(article).Association("Tags").Replace(tags)
}

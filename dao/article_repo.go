package dao

import (
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/models/res"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
func (a *ArticleRepo) GetList(pageSize, pageNum int) ([]res.Article, error) {
	var articles []res.Article
	err := a.db.
		Preload(clause.Associations).                     // 预加载全部
		Select(res.ArticleListClumns).                    // 指定查询字段
		Offset((pageNum - 1) * pageSize).Limit(pageSize). // 分页
		Order("created_at desc").Find(&articles).Error    // 排序查询
	if err != nil {
		return nil, err
	}
	return articles, err
}

// GetByUuid 根据uuid获取文章
func (a *ArticleRepo) GetByUuid(uuid string) (res.Article, error) {
	var article res.Article
	err := a.db.
		Preload(clause.Associations).
		Select(res.ArticleClumns).
		Where("uuid = ?", uuid).First(&article).Error
	if err != nil {
		return article, err
	}
	return article, err
}

// 根据uuid删除文章
func (a *ArticleRepo) DeleteByUuid(uuid string) error {
	tx := a.db.Begin()
	var article entity.Article
	if err := tx.Preload("Tags").Where("uuid = ?", uuid).First(&article).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除文章对应的标签
	if err := tx.Model(&article).Association("Tags").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	// 删除文章
	if err := tx.Delete(&article).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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

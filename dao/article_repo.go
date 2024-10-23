package dao

import (
	"encoding/json"
	"fmt"
	"gvb/app"
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/models/search"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticleRepo struct {
	db     *gorm.DB
	cache  *redis.Client
	search *meilisearch.Client
	index  string
}

var model = &entity.Article{}

func NewArticleRepo(db *gorm.DB, cache *redis.Client, search *meilisearch.Client, index string) *ArticleRepo {
	return &ArticleRepo{db: db, cache: cache, search: search, index: index}
}

func (a *ArticleRepo) GetListClumns() []string {
	return []string{"id", "created_at", "updated_at", "uuid", "title", "summary", "cover_image", "category_id", "views"}
}

func (a *ArticleRepo) GetManageListClumns() []string {
	return []string{"id", "created_at", "updated_at", "uuid", "title", "summary", "cover_image", "category_id", "views", "status", "top"}
}

func (a *ArticleRepo) GetAllClumns() []string {
	return []string{"id", "created_at", "updated_at", "uuid", "title", "summary", "content", "cover_image", "category_id", "views", "status", "top"}
}

func (a *ArticleRepo) Create(article *entity.Article) error {
	return a.db.Create(article).Error
}

// GetListOption 根据条件获取文章列表
func (a *ArticleRepo) GetListOption(isManage bool, pageSize, pageNum int, titleLike string, cateID int) ([]res.Article, int64, error) {
	var articles []res.Article
	var count int64
	query := a.db.Model(model).Preload(clause.Associations)
	if isManage {
		query = query.Select(a.GetManageListClumns())
	} else {
		query = query.Select(a.GetListClumns()).Where("status = ?", 1)
	}

	if titleLike != "" {
		query = query.Where("title like ?", fmt.Sprintf("%s%s%s", "%", titleLike, "%"))
	}

	if cateID != 0 {
		query = query.Where("category_id = ?", cateID)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Scopes(Paginate(pageNum, pageSize), OrderCreatedAtDesc).Find(&articles).Error; err != nil {
		return nil, 0, err
	}
	return articles, count, nil
}

// GetByUuid 根据uuid获取文章
func (a *ArticleRepo) GetByUuid(uuid string) (*res.Article, error) {
	article := new(res.Article)
	err := a.db.
		Preload(clause.Associations).
		Select(a.GetAllClumns()).
		Where("uuid = ?", uuid).First(article).Error
	if err != nil {
		return nil, err
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

	// 删除文章在搜索引擎中的记录
	if err := a.DeleteSearch(uuid); err != nil {
		return err
	}

	return tx.Commit().Error
}

// 根据uuid更新文章
func (a *ArticleRepo) UpdateByUuid(newArticle *req.Article, uuid uint32) (*entity.Article, error) {
	tx := a.db.Begin()
	article := new(entity.Article)
	if err := tx.Where("uuid = ?", uuid).First(article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	article.Title = newArticle.Title
	article.Content = newArticle.Content
	article.CoverImage = newArticle.CoverImage
	article.CategoryID = newArticle.CategoryID
	article.Top = newArticle.Top
	article.Status = newArticle.Status
	article.CreatedAt = time.UnixMilli(newArticle.CreatedAt)

	if err := tx.Save(article).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return article, tx.Commit().Error
}

// 更新文章的某个字段
func (a *ArticleRepo) UpdateSectionByUuid(uuid uint32, key string, value any) error {
	return a.db.Model(model).Where("uuid = ?", uuid).Update(key, value).Error
}

// 更新文章的标签
func (a *ArticleRepo) UpdateTags(article *entity.Article, tags []entity.Tag) error {
	return a.db.Model(article).Association("Tags").Replace(tags)
}

// 添加文章到搜索引擎
func (a *ArticleRepo) AddToSearch(article *entity.Article) {
	articleSearch := search.ArticleSearch{
		Uuid:    article.Uuid,
		Title:   article.Title,
		Summary: article.Summary,
		Content: article.Content,
	}
	_, err := a.search.Index(a.index).AddDocuments(articleSearch)
	if err != nil {
		app.Log.Warnf("Article:%d insert to meilisearch err:%v", articleSearch.Uuid, err)
	}
}

// 删除文章在搜索引擎中的记录
func (a *ArticleRepo) DeleteSearch(uuid string) error {
	_, err := a.search.Index(a.index).DeleteDocument(uuid)
	if err != nil {
		return err
	}
	return nil
}

// 更新搜索引擎中文章的摘要
func (a *ArticleRepo) UpdateSummarySearch(uuid uint32, summary string) (err error) {
	article := map[string]any{
		"uuid":    uuid,
		"summary": summary,
	}

	_, err = a.search.Index(a.index).UpdateDocuments(article)
	if err != nil {
		return err
	}

	return nil
}

// 搜索文章 搜索引擎实现
func (a *ArticleRepo) GetSearchList(query string) (*search.ArticleSearchResult, error) {
	res, err := a.search.Index(a.index).Search(query, &meilisearch.SearchRequest{
		AttributesToCrop:      []string{"title", "content", "summary"},
		CropLength:            30,
		AttributesToHighlight: []string{"title", "content", "summary"},
		HighlightPreTag:       "<span class=\"highlight\">",
		HighlightPostTag:      "</span>",
		MatchingStrategy:      "frequency",
	})
	if err != nil {
		return nil, err
	}

	byteres, err := json.Marshal(res)
	var newres search.SearchResponse
	if err := json.Unmarshal(byteres, &newres); err != nil {
		return nil, err
	}

	articles := new(search.ArticleSearchResult)
	for _, hit := range newres.Hits {
		articles.Hits = append(articles.Hits, search.ArticleSearch{
			Uuid:    hit.UUID,
			Title:   hit.Formatted.Title,
			Summary: hit.Formatted.Summary,
			Content: hit.Formatted.Content,
		})
	}

	articles.ProcessingTimeMs = res.ProcessingTimeMs
	articles.Total = res.EstimatedTotalHits

	return articles, nil
}

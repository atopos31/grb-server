package service

import (
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/models/res"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ArticleService struct {
	articleRepo *dao.ArticleRepo
	tagRepo     *dao.TagRepo
}

func NewArticleService(repo *dao.ArticleRepo, tagRepo *dao.TagRepo) *ArticleService {
	return &ArticleService{articleRepo: repo, tagRepo: tagRepo}
}

func (a *ArticleService) Create(reqArticle *req.Article) (uint32, error) {
	tags, err := a.tagRepo.FirstOrCreateTags(reqArticle.TagNames)
	if err != nil {
		return 0, err
	}
	article := entity.Article{
		Title:      reqArticle.Title,
		Uuid:       uuid.New().ID(),
		Content:    reqArticle.Content,
		CoverImage: reqArticle.CoverImage,
		CategoryID: reqArticle.CategoryID,
		Top:        reqArticle.Top,
		Status:     reqArticle.Status,
		Tags:       tags,
	}

	//自定义发布时间
	if reqArticle.CreatedAt != "" {
		int64time, _ := strconv.ParseInt(reqArticle.CreatedAt, 10, 64)
		article.CreatedAt = time.UnixMilli(int64time)
	}

	if err := a.articleRepo.Create(article); err != nil {
		return 0, err
	}

	return article.Uuid, nil
}

func (a *ArticleService) GetList(reqPage *req.ArticleList) ([]res.Article, error) {
	return a.articleRepo.GetList(reqPage.PageSize, reqPage.PageNum)
}

func (a *ArticleService) GetByUuid(uuid string) (res.Article, error) {
	return a.articleRepo.GetByUuid(uuid)
}

func (a *ArticleService) DeleteByUuid(uuid string) error {
	return a.articleRepo.DeleteByUuid(uuid)
}

func (a *ArticleService) Update(newArticle *req.Article, uuid string) error {
	article, err := a.articleRepo.UpdateByUuid(newArticle, uuid)
	if err != nil {
		return err
	}

	tags, err := a.tagRepo.FirstOrCreateTags(newArticle.TagNames)
	if err != nil {
		return err
	}

	return a.articleRepo.UpdateTags(article, tags)
}

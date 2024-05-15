package service

import (
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"

	"github.com/google/uuid"
)

type ArticleService struct {
	articleRepo *dao.ArticleRepo
	tagRepo     *dao.TagRepo
}

func NewArticleService(repo *dao.ArticleRepo, tagRepo *dao.TagRepo) *ArticleService {
	return &ArticleService{articleRepo: repo, tagRepo: tagRepo}
}

func (a *ArticleService) Create(reqArticle *req.Article) error {
	tags, err := a.tagRepo.FirstOrCreateTags(reqArticle.Tags)
	if err != nil {
		return err
	}
	article := entity.Article{
		Title:      reqArticle.Title,
		Uuid:       uuid.NewString(),
		Content:    reqArticle.Content,
		Summary:    reqArticle.Summary,
		CoverImage: reqArticle.CoverImage,
		CategoryID: reqArticle.CategoryID,
		Top:        reqArticle.Top,
		Status:     reqArticle.Status,
		Tags:       tags,
	}

	if err := a.articleRepo.Create(article); err != nil {
		return err
	}

	return nil
}

func (a *ArticleService) GetList(reqPage *req.ArticleList) ([]entity.Article, error) {
	return a.articleRepo.GetList(reqPage.PageSize, reqPage.PageNum)
}

func (a *ArticleService) GetByUuid(uuid string) (entity.Article, error) {
	return a.articleRepo.GetByUuid(uuid)
}

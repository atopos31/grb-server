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
	tags, err := a.tagRepo.FirstOrCreateTags(reqArticle.TagNames)
	if err != nil {
		return err
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

func (a *ArticleService) DeleteByUuid(uuid string) error {
	return a.articleRepo.DeleteByUuid(uuid)
}

func (a *ArticleService) Update(newArticle *req.Article) error {
	article, err := a.articleRepo.UpdateByUuid(newArticle)
	if err != nil {
		return err
	}

	tags, err := a.tagRepo.FirstOrCreateTags(newArticle.TagNames)
	if err != nil {
		return err
	}

	return a.articleRepo.UpdateTags(article, tags)
}
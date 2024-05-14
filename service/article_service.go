package service

import (
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"
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

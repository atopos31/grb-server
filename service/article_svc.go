package service

import (
	"gvb/dao"
	"gvb/global"
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/models/res"
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

func (a *ArticleService) Create(reqArticle *req.Article) (*res.ArticleCreateOrUpdate, error) {
	tags, err := a.tagRepo.FirstOrCreateTags(reqArticle.TagNames)
	if err != nil {
		return nil, err
	}

	article := entity.Article{
		Title:      reqArticle.Title,
		Uuid:       uuid.New().ID(), // 随机生成uuid
		Content:    reqArticle.Content,
		CoverImage: reqArticle.CoverImage,
		CategoryID: reqArticle.CategoryID,
		Top:        reqArticle.Top,
		Status:     reqArticle.Status,
		Tags:       tags,
	}
	article.CreatedAt = time.UnixMilli(reqArticle.CreatedAt)
	// 默认文章简介
	if len([]rune(article.Content)) > 200 {
		article.Summary = string([]rune(article.Content)[:200])
		global.Log.Info(article.Summary)
	} else {
		article.Summary = article.Content
	}

	if err := a.articleRepo.Create(article); err != nil {
		return nil, err
	}
	return &res.ArticleCreateOrUpdate{Uuid: article.Uuid, Status: article.Status}, nil
}

func (a *ArticleService) GetList(reqPage *req.ArticleList) (*res.ArticleList, error) {
	list, count, err := a.articleRepo.GetListOption(reqPage.PageSize, reqPage.PageNum, false, reqPage.TitleLike, reqPage.CategoryID)
	if err != nil {
		return nil, err
	}
	return &res.ArticleList{Count: count, List: list}, nil
}

func (a *ArticleService) GetManageList(reqPage *req.ArticleList) (*res.ArticleList, error) {
	list, count, err := a.articleRepo.GetListOption(reqPage.PageSize, reqPage.PageNum, true, reqPage.TitleLike, reqPage.CategoryID)
	if err != nil {
		return nil, err
	}
	return &res.ArticleList{Count: count, List: list}, nil
}

func (a *ArticleService) GetByUuid(uuid string) (*res.Article, error) {
	return a.articleRepo.GetByUuid(uuid)
}

func (a *ArticleService) DeleteByUuid(uuid string) error {
	return a.articleRepo.DeleteByUuid(uuid)
}

func (a *ArticleService) Update(newArticle *req.Article, uuid string) (*res.ArticleCreateOrUpdate, error) {
	article, err := a.articleRepo.UpdateByUuid(newArticle, uuid)
	if err != nil {
		return nil, err
	}

	tags, err := a.tagRepo.FirstOrCreateTags(newArticle.TagNames)
	if err != nil {
		return nil, err
	}

	if err := a.articleRepo.UpdateTags(article, tags); err != nil {
		return nil, err
	}

	return &res.ArticleCreateOrUpdate{Uuid: article.Uuid, Status: article.Status}, nil
}

func (a *ArticleService) UpdateSectionByUuid(uuid string, section *req.ArticleSertion) error {
	return a.articleRepo.UpdateSectionByUuid(uuid, section.Key, section.Value)
}

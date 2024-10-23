package service

import (
	"gvb/app"
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/models/search"
	"time"
)

type ArticleService struct {
	articleRepo *dao.ArticleRepo
	tagRepo     *dao.TagRepo
}

func NewArticleService(repo *dao.ArticleRepo, tagRepo *dao.TagRepo) *ArticleService {
	return &ArticleService{articleRepo: repo, tagRepo: tagRepo}
}

func (a *ArticleService) Create(reqArticle *req.Article, uuid uint32) (*res.ArticleCreateOrUpdate, error) {
	tags, err := a.tagRepo.FirstOrCreateTags(reqArticle.TagNames)
	if err != nil {
		return nil, err
	}

	article := &entity.Article{
		Title:      reqArticle.Title,
		Uuid:       uuid,
		Content:    reqArticle.Content,
		CoverImage: reqArticle.CoverImage,
		CategoryID: reqArticle.CategoryID,
		Top:        reqArticle.Top,
		Status:     reqArticle.Status,
		Tags:       tags,
	}
	article.CreatedAt = time.UnixMilli(reqArticle.CreatedAt)
	// 默认文章简介
	contentRune := []rune(article.Content) // 将文章内容转为rune切片
	if len(contentRune) > 200 {
		article.Summary = string(contentRune[:200])
	} else {
		article.Summary = article.Content
	}

	if err := a.articleRepo.Create(article); err != nil {
		return nil, err
	}

	// 更新文章到搜索
	go a.articleRepo.AddToSearch(article)

	return &res.ArticleCreateOrUpdate{Uuid: article.Uuid, Status: article.Status}, nil
}

func (a *ArticleService) GetList(reqPage *req.ArticleList) (*res.ArticleList, error) {
	list, count, err := a.articleRepo.GetListOption(false, reqPage.PageSize, reqPage.PageNum, reqPage.TitleLike, reqPage.CategoryID)
	if err != nil {
		return nil, err
	}
	return &res.ArticleList{Count: count, List: list}, nil
}

func (a *ArticleService) GetManageList(reqPage *req.ArticleList) (*res.ArticleList, error) {
	list, count, err := a.articleRepo.GetListOption(true, reqPage.PageSize, reqPage.PageNum, reqPage.TitleLike, reqPage.CategoryID)
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

func (a *ArticleService) Update(newArticle *req.Article, uuid uint32) (*res.ArticleCreateOrUpdate, error) {
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

	// 更新文章到搜索
	go a.articleRepo.AddToSearch(article)

	return &res.ArticleCreateOrUpdate{Uuid: article.Uuid, Status: article.Status}, nil
}

// 可选字段更新
func (a *ArticleService) UpdateSectionByUuid(uuid uint32, section *req.ArticleSertion) error {
	return a.articleRepo.UpdateSectionByUuid(uuid, section.Key, section.Value)
}

// 从搜索引擎搜索
func (a *ArticleService) Search(query string) (*search.ArticleSearchResult, error) {
	return a.articleRepo.GetSearchList(query)
}

// 生成摘要并加载到数据库和搜索引擎
func (a *ArticleService) GenerateSummary(uuid uint32, content string) {
	summary, err := Svc.AiService.GetSummary(content)
	if err != nil {
		app.Log.Warnf("Article generate summary err:%v uuid:%d", err, uuid)
		return
	}

	if err := a.articleRepo.UpdateSectionByUuid(uuid, "summary", summary); err != nil {
		app.Log.Warnf("Article gpdate summary err:%v uuid:%d", err, uuid)
		return
	}
	if err := a.articleRepo.UpdateSummarySearch(uuid, summary); err != nil {
		app.Log.Warnf("Article update summary to meilisearch  err:%v uuid:%d", err, uuid)
		return
	}
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/google/wire"
	"gvb/config"
	"gvb/core"
	"gvb/dao"
)

// Injectors from wire.go:

func New(config2 config.Config) (*Service, error) {
	mysql := config.NewMysqlConfig(config2)
	db := core.NewGormDB(mysql)
	userRepo := dao.NewUserRepo(db)
	userService := NewUserService(userRepo)
	redis := config.NewRedisConfig(config2)
	client := core.NewRedisCache(redis)
	meilisearch := config.NewMeilisearchConfig(config2)
	meilisearchClient := core.NewMeiliSearchClient(meilisearch)
	string2 := config.NewSiteInfoPath(config2)
	articleRepo := dao.NewArticleRepo(db, client, meilisearchClient, string2)
	tagRepo := dao.NewTagRepo(db)
	articleService := NewArticleService(articleRepo, tagRepo)
	cateRepo := dao.NewCateRepo(db)
	cateService := NewCateService(cateRepo)
	tagService := NewTagService(tagRepo)
	commentRepo := dao.NewCommentRepo(db)
	commentService := NewCommentService(commentRepo)
	siteInfoService := NewSiteInfoService(db, client, string2)
	ossQiniu := config.NewOssQiniuConfig(config2)
	ossService := NewOssQinui(ossQiniu)
	ai := config.NewAiConfig(config2)
	aiService := NewAiSvcByConfig(ai)
	service := NewService(config2, userService, articleService, cateService, tagService, commentService, siteInfoService, ossService, aiService)
	return service, nil
}

// wire.go:

var SvcSet = wire.NewSet(config.ProvideSet, core.ProvideSet, dao.ProvideSet, ProvideSet)

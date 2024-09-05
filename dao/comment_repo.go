package dao

import (
	"gvb/models/entity"
	"gvb/models/res"

	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

func (c *CommentRepo) Create(comment *entity.Comment) error {
	return c.db.Create(comment).Error
}

func (c *CommentRepo) GetCommentList(status uint8) ([]*res.CommentManager, error) {
	var Comments []*res.CommentManager
	if err := c.db.Model(&entity.Comment{}).
		Preload("Article", func(db *gorm.DB) *gorm.DB { return db.Select("title", "uuid") }).
		Where("status = ?", status).
		Order("created_at desc").
		Find(&Comments).Error; err != nil {

		return nil, err
	}

	for _, comment := range Comments {
		comment.ArticleTitle = comment.Article.Title
	}

	return Comments, nil
}

func (c *CommentRepo) GetArticleCommentListByUuid(uuid uint32) ([]*res.Comment, error) {
	var rootComments []*res.Comment
	if err := c.db.Model(&entity.Comment{}).
		Where("status = ?", 1).
		Where("article_uuid = ?", uuid).
		Where("root_id is null").
		Order("created_at desc").Find(&rootComments).Error; err != nil {
		return nil, err
	}

	for _, rootComment := range rootComments {
		if err := c.db.Model(&entity.Comment{}).Where("root_id = ?", rootComment.ID).Find(&rootComment.ChildComment).Error; err != nil {
			return nil, err
		}
	}

	return rootComments, nil
}

func (c *CommentRepo) UpdateStatus(id uint, status uint8) error {
	return c.db.Model(&entity.Comment{}).Where("id = ?", id).Update("status", status).Error
}

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

func (c *CommentRepo) GetList(uuid uint32) ([]*res.Comment, error) {
	var rootComments []*res.Comment
	if err := c.db.Model(&entity.Comment{}).Where("article_uuid = ?", uuid).Where("root_id is null").Find(&rootComments).Error; err != nil {
		return nil, err
	}

	for _, rootComment := range rootComments {
		if err := c.db.Model(&entity.Comment{}).Where("root_id = ?", rootComment.ID).Find(&rootComment.ChildComment).Error; err != nil {
			return nil, err
		}
	}

	return rootComments, nil
}

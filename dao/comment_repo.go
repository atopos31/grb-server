package dao

import (
	"gvb/models/entity"

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

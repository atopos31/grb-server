package service

import (
	"fmt"
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/utils/conver"
)

type CommentService struct {
	CommentRepo *dao.CommentRepo
}

func NewCommentService(commentRepo *dao.CommentRepo) *CommentService {
	return &CommentService{
		CommentRepo: commentRepo,
	}
}

const gravatarCNURl = "https://www.gravatar.cn/avatar/"
const gravatarCOMURl = "https://www.gravatar.com/avatar/"

func (c *CommentService) Create(reqComment *req.Comment) error {
	emailMD5 := conver.StringToMD5(reqComment.Email)

	comment := &entity.Comment{
		ArticleUuid: reqComment.ArticleID,
		UserName:    reqComment.UserName,
		Avatar:      fmt.Sprintf("%s%s", gravatarCNURl, emailMD5),
		Email:       reqComment.Email,
		Content:     reqComment.Content,
		ParentID:    reqComment.ParentID,
		RootID:      reqComment.RootID,
	}
	return c.CommentRepo.Create(comment)
}

package service

import (
	"fmt"
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/models/res"
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

// 头像地址
const cravatarURl = "https://cravatar.cn/avatar/"
const gravatarURl = "https://gravatar.com/avatar/"

func (c *CommentService) Create(reqComment *req.Comment) error {
	emailMD5 := conver.StringToMD5(reqComment.Email)

	comment := &entity.Comment{
		ArticleUuid: reqComment.ArticleID,
		UserName:    reqComment.UserName,
		Avatar:      fmt.Sprintf("%s%s", cravatarURl, emailMD5),
		Email:       reqComment.Email,
		WebSite:     reqComment.WebSite,
		Content:     reqComment.Content,
		ParentID:    reqComment.ParentID,
		RootID:      reqComment.RootID,
	}
	return c.CommentRepo.Create(comment)
}

func (c *CommentService) GetList(uuid uint32) ([]*res.Comment, error) {
	return c.CommentRepo.GetList(uuid)
}

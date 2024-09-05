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
	avatarURL := fmt.Sprintf("%s%s", cravatarURl, emailMD5)

	comment := &entity.Comment{
		ArticleUuid: reqComment.ArticleID,
		UserName:    reqComment.UserName,
		Avatar:      avatarURL,
		Email:       reqComment.Email,
		WebSite:     reqComment.WebSite,
		Content:     reqComment.Content,
		ParentID:    reqComment.ParentID,
		RootID:      reqComment.RootID,
	}
	return c.CommentRepo.Create(comment)
}

func (c *CommentService) GetArticleCommentListByUuid(uuid uint32) ([]*res.Comment, error) {
	return c.CommentRepo.GetArticleCommentListByUuid(uuid)
}

func (c *CommentService) GetCommentList(status uint8) ([]*res.CommentManager, error) {
	return c.CommentRepo.GetCommentList(status)
}

func (c *CommentService) UpdateStatus(id uint, status uint8) error {
	return c.CommentRepo.UpdateStatus(id, status)
}

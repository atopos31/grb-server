package entity

import (
	"gorm.io/gorm"
)

// 文章表
type Article struct {
	gorm.Model
	Uuid       uint32 `gorm:"not null;index;comment:文章uuid;unique"`
	Title      string `gorm:"type:varchar(255);not null;comment:文章标题"`
	Summary    string `gorm:"type:text;comment:文章摘要"`
	Content    string `gorm:"type:longtext;not null;comment:文章内容"`
	CoverImage string `gorm:"type:varchar(255);comment:文章封面"`
	Top        uint8  `gorm:"type:int(1);default:0;comment:是否置顶"`
	Status     uint8  `gorm:"type:int(1);default:0;comment:文章状态"`
	Views      int    `gorm:"type:int(11);default:0;comment:文章浏览量"`
	CategoryID uint   `gorm:"not null;comment:文章分类"`
	Category   Category
	Tags       []Tag `gorm:"many2many:articles_tags;"`
}

// 文章分类
type Category struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255);not null"`
	Articles []Article `gorm:"foreignKey:CategoryID"`
}

// 文章标签
type Tag struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255);not null"`
	Articles []Article `gorm:"many2many:articles_tags;"`
}

// 评论表
type Comment struct {
	gorm.Model
	ArticleUuid uint32   `gorm:"not null;index;comment:所属文章ID"`
	Article     Article  `gorm:"foreignKey:ArticleUuid;references:Uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserName    string   `gorm:"not null;size:255;comment:评论用户名"`
	Email       string   `gorm:"not null;size:255;comment:评论邮箱"`
	Avatar      string   `gorm:"not null;size:255;comment:评论头像"`
	Status      uint8    `gorm:"type:int(1);default:0;comment:评论状态 0 待审核 1 审核通过 2 审核不通过"`
	RootID      *uint    `gorm:"index;default:NULL;comment:根评论ID，为NULL说明是一级评论"`
	ParentID    *uint    `gorm:"index;default:NULL;comment:父评论ID，在RootID不为NULL的情况下，为NULL说明是二级评论，不为NUll说明是二级评论的回复评论"`
	Root        *Comment `gorm:"foreignKey:RootID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Parent      *Comment `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WebSite     string   `gorm:"size:255;comment:访客的网站"`
	Content     string   `gorm:"not null;size:1024;comment:评论内容"`
}

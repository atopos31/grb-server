package entity

import (
	"gorm.io/gorm"
)

// 文章表
type Article struct {
	gorm.Model
	Uuid       uint32 `gorm:"not null;index;comment:文章uuid;unique"`
	Title      string `gorm:"type:varchar(255);not null;comment:文章标题"`
	Summary    string `gorm:"type:varchar(255);comment:文章摘要"`
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
	Name     string    `gorm:"type:varchar(255);not null;unique"`
	Articles []Article `gorm:"foreignKey:CategoryID"`
}

// 文章标签
type Tag struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255);not null;unique"`
	Articles []Article `gorm:"many2many:articles_tags;"`
}

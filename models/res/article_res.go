package res

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Uuid       uint32
	Title      string
	Summary    string
	Content    string
	CoverImage string
	Views      int
	CategoryID uint
	Category   Category
	Tags       []Tag          `gorm:"many2many:articles_tags;"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"` //用来避免查询到软删除的数据
}

var ArticleListClumns = []string{"id", "created_at", "updated_at", "uuid", "title", "summary", "cover_image", "category_id", "views"}
var ArticleClumns = []string{"id", "created_at", "updated_at", "uuid", "title", "summary", "content", "cover_image", "category_id", "views"}

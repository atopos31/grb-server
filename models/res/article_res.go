package res

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	Uuid       uint32         `json:"uuid"`
	Title      string         `json:"title"`
	Summary    string         `json:"summary"`
	Content    string         `json:"content"`
	CoverImage string         `json:"cover_image"`
	Views      int            `json:"views"`
	Status     uint8          `json:"status"`
	Top        uint8          `json:"top"`
	CategoryID uint           `json:"category_id"`
	Category   Category       `json:"category"`
	Tags       []Tag          `gorm:"many2many:articles_tags;" json:"tags"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"` //用来避免查询到软删除的数据
}

type ArticleCreateOrUpdate struct {
	Uuid   uint32 `json:"uuid"`
	Status uint8  `json:"status"`
}

type ArticleList struct {
	Count int64     `json:"count"`
	List  []Article `json:"list"`
}

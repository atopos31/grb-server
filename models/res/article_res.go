package res

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type LocalTime time.Time

type Article struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	CreatedAt  LocalTime      `json:"created_at"`
	UpdatedAt  LocalTime      `json:"updated_at"`
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

// 序列化成时间戳
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return json.Marshal(tTime.UnixMilli())
}

type ArticleCreateOrUpdate struct {
	Uuid   uint32 `json:"uuid"`
	Status uint8  `json:"status"`
}

type ArticleList struct {
	Count int64     `json:"count"`
	List  []Article `json:"list"`
}

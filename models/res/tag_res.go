package res

import "gorm.io/gorm"

type Tag struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` //用来避免查询到软删除的数据
}

type ManageTag struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Count     int            `json:"count"`
	CreatedAt int64          `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` //用来避免查询到软删除的数据
}

type ManageTagList struct {
	Count int64       `json:"count"`
	List  []ManageTag `json:"list"`
}

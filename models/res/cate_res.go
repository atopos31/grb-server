package res

import "gorm.io/gorm"

type Category struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` //用来避免查询到软删除的数据
}

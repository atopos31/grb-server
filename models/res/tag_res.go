package res

import "gorm.io/gorm"

type Tag struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` //用来避免查询到软删除的数据
}

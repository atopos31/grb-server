package entity

import (
	"gorm.io/gorm"
)

func InitEntity(db *gorm.DB) {
	// 创建表
	err := db.AutoMigrate(&AdminUser{}, &Category{}, &Article{}, &Tag{}, &Comment{})
	if err != nil {
		panic(err)
	}
}

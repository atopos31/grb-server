package dao

import "gorm.io/gorm"

// 分页
func Paginate(num, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if num <= 0 {
			num = 1
		}
		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}
		offset := (num - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

// gorm.model 的默认字段 以创建时间倒序
func OrderCreatedAtDesc(db *gorm.DB) *gorm.DB {
	return db.Order("created_at desc")
}

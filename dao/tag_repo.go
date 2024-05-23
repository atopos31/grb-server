package dao

import (
	"gvb/models/entity"

	"gorm.io/gorm"
)

type TagRepo struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{db: db}
}

/*
在插入文章之前
先查询数据库有没有这个标签
如果有就返回这个标签
如果没有就创建这个标签
并返回完整可关联的数据
*/
func (t *TagRepo) FirstOrCreateTags(tagnames []string) ([]entity.Tag, error) {
	var realTags []entity.Tag

	tx := t.db.Begin()
	for _, tagname := range tagnames {
		tag := entity.Tag{Name: tagname}
		if err := tx.FirstOrCreate(&tag, tag).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		realTags = append(realTags, tag)
	}

	return realTags, tx.Commit().Error
}

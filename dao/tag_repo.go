package dao

import (
	"gvb/global"
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
func (t *TagRepo) FirstOrCreateTags(tags []string) ([]entity.Tag, error) {
	var realTags []entity.Tag
	for _, v := range tags {
		//go 1.22新增特性
		global.Log.Info(v)
		tag := entity.Tag{Name: v}
		if err := t.db.FirstOrCreate(&tag, tag).Error; err != nil {
			return nil, err
		}
		realTags = append(realTags, tag)
	}

	return realTags, nil
}

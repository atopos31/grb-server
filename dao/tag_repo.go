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

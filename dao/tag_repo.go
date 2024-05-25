package dao

import (
	"errors"
	"gvb/models/entity"
	"gvb/models/errcode"

	"github.com/go-sql-driver/mysql"
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

func (t *TagRepo) Create(tag entity.Tag) error {
	if err := t.db.Create(&tag).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// 1062 unique数据重复冲突
			return errcode.ErrDataIsExits
		} else {
			return err
		}
	}

	return nil
}

func (t *TagRepo) GetList() ([]entity.Tag, error) {
	var tags []entity.Tag
	if err := t.db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *TagRepo) Update(name string, id uint) error {
	if err := t.db.Model(&entity.Tag{}).Where("id = ?", id).Update("name", name).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// 1062 unique数据重复冲突
			return errcode.ErrDataIsExits
		} else {
			return err
		}
	}

	return nil
}

func (t *TagRepo) Delete(id uint) error {
	return t.db.Delete(&entity.Tag{}, id).Error
}

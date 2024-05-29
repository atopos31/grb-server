package dao

import (
	"errors"
	"gvb/models/entity"
	"gvb/models/errcode"
	"gvb/models/res"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CateRepo struct {
	db *gorm.DB
}

func NewCateRepo(db *gorm.DB) *CateRepo {
	return &CateRepo{db: db}
}

func (c *CateRepo) Create(cate entity.Category) (uint, error) {
	if err := c.db.Create(&cate).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			// 1062 unique数据重复冲突
			return 0, errcode.ErrDataIsExits
		} else {
			return 0, err
		}
	}

	return cate.ID, nil
}

func (c *CateRepo) GetList() ([]res.Category, error) {
	var cates []res.Category
	err := c.db.Select("id", "name").Order("created_at asc").Find(&cates).Error
	return cates, err
}

func (c *CateRepo) GetByName(name string) (entity.Category, error) {
	var cate entity.Category
	err := c.db.Where("name = ?", name).First(&cate).Error
	return cate, err
}

func (c *CateRepo) GetByID(id uint) (entity.Category, error) {
	var cate entity.Category
	err := c.db.Where("id = ?", id).First(&cate).Error
	return cate, err
}

func (c *CateRepo) Update(name string, id uint) error {
	if err := c.db.Model(&entity.Category{}).Where("id = ?", id).Update("name", name).Error; err != nil {
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

func (c *CateRepo) Delete(id uint) error {
	return c.db.Delete(&entity.Category{}, id).Error
}

package dao

import (
	"errors"
	"gvb/models/entity"
	"gvb/models/errcode"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CateRepo struct {
	db *gorm.DB
}

func NewCateRepo(db *gorm.DB) *CateRepo {
	return &CateRepo{db: db}
}

func (c *CateRepo) Create(cate entity.Category) error {
	err := c.db.Create(&cate).Error
	if err != nil {
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

func (c *CateRepo) GetList() ([]entity.Category, error) {
	var cates []entity.Category
	err := c.db.Find(&cates).Error
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

func (c *CateRepo) Update(cate entity.Category) error {
	return c.db.Save(&cate).Error
}

func (c *CateRepo) Delete(id uint) error {
	return c.db.Delete(&entity.Category{}, id).Error
}

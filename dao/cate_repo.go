package dao

import (
	"gvb/models/entity"
	"gvb/models/errcode"
	"gvb/models/res"

	"gorm.io/gorm"
)

type CateRepo struct {
	db *gorm.DB
}

func NewCateRepo(db *gorm.DB) *CateRepo {
	return &CateRepo{db: db}
}

func (c *CateRepo) Create(cate entity.Category) (uint, error) {
	if err := c.db.Model(&entity.Category{}).Create(&cate).Error; err != nil {
		if errcode.CheckMysqlErrDataIsExits(err) {
			return 0, errcode.ErrDataIsExits
		} else {
			return 0, err
		}
	}

	return cate.ID, nil
}

func (c *CateRepo) GetList() ([]res.Category, error) {
	var cates []res.Category
	err := c.db.Select("id", "name").Scopes(OrderCreatedAtDesc).Find(&cates).Error
	return cates, err
}

func (c *CateRepo) GetManageList(pageSize, pageNum int) ([]res.ManageCategory, error) {
	var cates []entity.Category
	err := c.db.Select("id", "name", "created_at").
		Preload("Articles").
		Scopes(Paginate(pageNum, pageSize), OrderCreatedAtDesc).
		Order("created_at desc").Find(&cates).Error
	if err != nil {
		return nil, err
	}
	var manageCategories []res.ManageCategory
	for _, cate := range cates {
		manageCategories = append(manageCategories, res.ManageCategory{
			Name:      cate.Name,
			Count:     len(cate.Articles),
			ID:        cate.ID,
			CreatedAt: cate.CreatedAt.UnixMilli(),
		})
	}
	return manageCategories, nil
}

func (c *CateRepo) GetCount() (int64, error) {
	var count int64
	err := c.db.Model(&entity.Category{}).Count(&count).Error
	return count, err
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
		if errcode.CheckMysqlErrDataIsExits(err) {
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

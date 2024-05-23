package service

import (
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"

	"gorm.io/gorm"
)

type CateService struct {
	cateRepo *dao.CateRepo
}

func NewCateService(repo *dao.CateRepo) *CateService {
	return &CateService{cateRepo: repo}
}

func (c *CateService) Create(reqCate *req.Cate) error {
	cate := entity.Category{
		Name: reqCate.Name,
	}
	return c.cateRepo.Create(cate)
}

func (c *CateService) GetList() ([]entity.Category, error) {
	return c.cateRepo.GetList()
}

func (c *CateService) Update(reqCate *req.Cate) error {
	cate := entity.Category{
		Model: gorm.Model{ID: reqCate.ID},
		Name:  reqCate.Name,
	}
	return c.cateRepo.Update(cate)
}

func (c *CateService) Delete(id uint) error {
	return c.cateRepo.Delete(id)
}

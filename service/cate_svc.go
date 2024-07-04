package service

import (
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/models/res"
)

type CateService struct {
	cateRepo *dao.CateRepo
}

func NewCateService(repo *dao.CateRepo) *CateService {
	return &CateService{cateRepo: repo}
}

func (c *CateService) Create(reqCate *req.Cate) (uint, error) {
	cate := entity.Category{
		Name: reqCate.Name,
	}
	return c.cateRepo.Create(cate)
}

func (c *CateService) GetList() ([]res.Category, error) {
	return c.cateRepo.GetList()
}

func (c *CateService) GetManageList(reqpage *req.CateList) (*res.ManageCategoryList, error) {
	count, err := c.cateRepo.GetCount()
	if err != nil {
		return nil, err
	}

	list, err := c.cateRepo.GetManageList(reqpage.PageSize, reqpage.PageNum)
	if err != nil {
		return nil, err
	}

	return &res.ManageCategoryList{Count: count, List: list}, nil
}

func (c *CateService) Update(reqCate *req.Cate) error {
	return c.cateRepo.Update(reqCate.Name, reqCate.ID)
}

func (c *CateService) Delete(id uint) error {
	return c.cateRepo.Delete(id)
}

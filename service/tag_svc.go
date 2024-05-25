package service

import (
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"
)

type TagService struct {
	tagRepo *dao.TagRepo
}

func NewTagService(repo *dao.TagRepo) *TagService {
	return &TagService{
		tagRepo: repo,
	}
}

func (t *TagService) Create(reqTag *req.Tag) error {
	tag := entity.Tag{
		Name: reqTag.Name,
	}

	return t.tagRepo.Create(tag)
}

func (t *TagService) GetList() ([]entity.Tag, error) {
	return t.tagRepo.GetList()
}

func (t *TagService) Update(reqTag *req.Tag) error {
	return t.tagRepo.Update(reqTag.Name, reqTag.ID)
}

func (t *TagService) Delete(id uint) error {
	return t.tagRepo.Delete(id)
}

package service

import (
	"gvb/dao"
	"gvb/models/entity"
	"gvb/models/req"
	"gvb/models/res"
)

type TagService struct {
	tagRepo *dao.TagRepo
}

func NewTagService(repo *dao.TagRepo) *TagService {
	return &TagService{
		tagRepo: repo,
	}
}

func (t *TagService) Create(reqTag *req.Tag) (uint, error) {
	tag := entity.Tag{
		Name: reqTag.Name,
	}

	return t.tagRepo.Create(tag)
}

func (t *TagService) GetList() ([]res.Tag, error) {
	return t.tagRepo.GetList()
}

func (t *TagService) GetManageList(reqpage *req.TagList) (*res.ManageTagList, error) {
	count, err := t.tagRepo.GetCount()
	if err != nil {
		return nil, err
	}

	list, err := t.tagRepo.GetManageList(reqpage.PageSize, reqpage.PageNum)
	if err != nil {
		return nil, err
	}

	return &res.ManageTagList{Count: count, List: list}, nil
}

func (t *TagService) GetHotList(size int) ([]res.Tag, error) {
	return t.tagRepo.GetHotList(size)
}

func (t *TagService) Update(reqTag *req.Tag) error {
	return t.tagRepo.Update(reqTag.Name, reqTag.ID)
}

func (t *TagService) Delete(id uint) error {
	return t.tagRepo.Delete(id)
}

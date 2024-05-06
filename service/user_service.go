package service

import (
	"gvb/dao"

	"gorm.io/gorm"
)

type UserService struct {
	repo *dao.UserRepo
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo: dao.NewUserRepo(db),
	}
}

func (u *UserService) Login() error {
	return nil
}

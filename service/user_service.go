package service

import (
	"gvb/dao"
)

type UserService struct {
	userrepo *dao.UserRepo
}

func NewUserService(repo *dao.UserRepo) *UserService {
	return &UserService{
		userrepo: repo,
	}
}

func (u *UserService) Login() error {
	return nil
}

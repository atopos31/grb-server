package service

import (
	"gvb/dao"
	"gvb/models/req"
	"gvb/models/res"
	"gvb/utils/jwt"
)

type UserService struct {
	userrepo *dao.UserRepo
}

func NewUserService(repo *dao.UserRepo) *UserService {
	return &UserService{
		userrepo: repo,
	}
}

func (u *UserService) Login(user *req.User) (string, error) {
	if err := u.userrepo.ValidateUser(user.Username, user.Password); err != nil {
		return "", err
	}

	token, err := jwt.Sign(jwt.Payload{
		Username: user.Username,
		Password: user.Password,
	})
	return token, err
}

func (u *UserService) GetInfo() (info *res.UserInfo, err error) {
	return u.userrepo.GetInfo()
}

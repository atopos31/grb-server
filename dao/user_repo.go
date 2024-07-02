package dao

import (
	"gvb/models/entity"
	"gvb/models/res"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) ValidateUser(username string, password string) error {
	var user entity.AdminUser
	return u.db.Where("name = ? and password = ?", username, password).First(&user).Error
}

func (u *UserRepo) GetInfo() (info *res.UserInfo, err error) {
	var user entity.AdminUser
	if err = u.db.Select("name", "avatar").First(&user).Error; err != nil {
		return nil, err
	}

	return &res.UserInfo{
		Username: user.Name,
		Avatar:   user.Avatar,
	}, nil
}

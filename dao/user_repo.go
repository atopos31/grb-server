package dao

import (
	"gvb/models/entity"

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

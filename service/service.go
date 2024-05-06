package service

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var Svc *Service

type Service struct {
	UserService *UserService
}

func New(db *gorm.DB, cache *redis.Client) *Service {
	return &Service{
		UserService: NewUserService(db),
	}
}

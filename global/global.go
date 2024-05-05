package global

import (
	"gvb/config"

	"gorm.io/gorm"
)

var (
	// Conf 全局配置
	Conf *config.Config
	DB   *gorm.DB
)

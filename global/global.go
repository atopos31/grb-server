package global

import (
	"gvb/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// Conf 全局配置
	Conf *config.Config
	// DB   数据库
	DB *gorm.DB
	// Log 日志
	Log *logrus.Logger
)

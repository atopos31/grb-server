package global

import (
	"gvb/config"

	"github.com/sirupsen/logrus"
)

var (
	// Conf 全局配置
	Conf *config.Config
	// Log 日志
	Log *logrus.Logger
)

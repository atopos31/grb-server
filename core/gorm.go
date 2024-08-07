package core

import (
	"errors"
	"time"

	"gvb/config"
	"gvb/models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库
func NewGormDB(config config.Mysql) *gorm.DB {
	if config.Host == "" || config.Port == "" {
		panic(errors.New("Mysql config error ,Host or port is NULL!"))
	}

	dsn := config.Dsn()

	logger := logger.Default
	loggerMode := configLevelToMode(config.LogLevel)
	logger.LogMode(loggerMode)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)              // 设置连接池
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 设置连接超时时间

	// 初始化表
	entity.InitEntity(db)

	return db
}

func configLevelToMode(logLevel string) logger.LogLevel {
	switch logLevel {
	case config.ENV_GORM_ERROR:
		return logger.Error
	case config.ENV_GORM_WORN:
		return logger.Warn
	case config.ENV_GORM_INFO:
		return logger.Info
	default:
		return logger.Info
	}
}

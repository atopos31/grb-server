package core

import (
	"gvb/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库
func InitGorm() *gorm.DB {
	if global.Conf.Mysql.Host == "" {
		return nil
	}

	dsn := global.Conf.Mysql.Dsn()

	// 配置日志
	var mysqllogger logger.Interface
	if global.Conf.Mysql.LogLevel == "info" {
		mysqllogger = logger.Default.LogMode(logger.Info)
	} else if global.Conf.Mysql.LogLevel == "error" {
		mysqllogger = logger.Default.LogMode(logger.Error)
	} else if global.Conf.Mysql.LogLevel == "warn" {
		mysqllogger = logger.Default.LogMode(logger.Warn)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqllogger,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)              // 	设置连接池
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 设置连接超时时间

	return db
}

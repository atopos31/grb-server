package core

import (
	"errors"
	"time"

	"gvb/config"
	"gvb/models/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库
func InitGorm(config config.Mysql) *gorm.DB {
	if config.Host == "" || config.Port == "" {
		panic(errors.New("mysql配置错误"))
	}

	dsn := config.Dsn()
	mysqllogger := config.Logger()

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

	// 初始化表
	entity.InitEntity(db)

	return db
}

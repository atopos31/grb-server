package config

import "gorm.io/gorm/logger"

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Config   string `mapstructure:"config"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	LogLevel string `mapstructure:"log_level"`
}

// Dsn mysql dsn
func (m *Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Config
}

// Logger mysql logger
func (m *Mysql) Logger() logger.Interface {
	var mysqllogger logger.Interface
	if m.LogLevel == "info" {
		mysqllogger = logger.Default.LogMode(logger.Info)
	} else if m.LogLevel == "error" {
		mysqllogger = logger.Default.LogMode(logger.Error)
	} else if m.LogLevel == "warn" {
		mysqllogger = logger.Default.LogMode(logger.Warn)
	}

	return mysqllogger
}

package config

import (
	"fmt"
)

const (
	ENV_GORM_INFO  = "info"
	ENV_GORM_ERROR = "error"
	ENV_GORM_WORN  = "warn"
)

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Config   string `mapstructure:"config"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	LogLevel string `mapstructure:"log_level"`
}

func NewMysqlConfig(config Config) Mysql {
	return config.Mysql
}

// Dsn mysql dsn
func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.User, m.Password, m.Host, m.Port, m.Database, m.Config)
}

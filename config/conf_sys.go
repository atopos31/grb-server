package config

import "fmt"

const (
	ENV_SYS_DEBUG   = "debug"
	ENV_SYS_RELEASE = "release"
	ENV_SYS_TEST    = "test"
)

type System struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Env          string `mapstructure:"env"`
	Origin       string `mapstructure:"origin"`
	SiteInfoPath string `mapstructure:"site_info_path"`
}

func NewSiteInfoPath(config Config) string {
	return config.Sys.SiteInfoPath
}

// 获取监听地址
func (s *System) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

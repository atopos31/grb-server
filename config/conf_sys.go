package config

type System struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Env          string `mapstructure:"env"`
	Origin       string `mapstructure:"origin"`
	SiteInfoPath string `mapstructure:"site_info_path"`
}

// 获取监听地址
func (s *System) Addr() string {
	return s.Host + ":" + s.Port
}

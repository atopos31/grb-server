package config

type System struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Env    string `mapstructure:"env"`
	Origin string `mapstructure:"origin"`
}

// addr
func (s *System) Addr() string {
	return s.Host + ":" + s.Port
}

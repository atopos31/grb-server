package config

type Jwt struct {
	Timeout int64  `mapstructure:"timeout"`
	Issuer  string `mapstructure:"issuer"`
	Secret  string `mapstructure:"secret"`
}

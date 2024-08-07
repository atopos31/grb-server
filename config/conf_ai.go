package config

type Ai struct {
	Use     string    `mapstructure:"use"`
	Hunyuan AiHunyuan `mapstructure:"hunyuan"`
	Qianfan AiQianfan `mapstructure:"qianfan"`
}

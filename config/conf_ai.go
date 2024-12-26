package config

type Ai struct {
	Use     string    `mapstructure:"use"`
	Hunyuan AiHunyuan `mapstructure:"hunyuan"`
	Qianfan AiQianfan `mapstructure:"qianfan"`
	Gemini  AiGemini  `mapstructure:"gemini"`
	Openai  AiOpenai  `mapstructure:"openai"`
}

func NewAiConfig(config Config) Ai {
	return config.Ai
}

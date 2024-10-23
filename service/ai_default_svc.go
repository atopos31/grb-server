package service

import (
	"fmt"
	"gvb/config"
)

type AiDefault struct {
	conf config.Ai
}

func NewAiDefault(conf config.Ai) *AiDefault {
	return &AiDefault{
		conf: conf,
	}
}

func (a *AiDefault) GetSummary(articleContent string) (string, error) {
	return fmt.Sprintf("AI服务未配置, 或参数错误%+v", a.conf), nil
}

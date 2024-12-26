package service

import (
	"context"
	"gvb/config"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type AiOpenai struct {
	client *openai.Client
	model  string
}

func NewAiOpenai(config config.AiOpenai) *AiOpenai {
	return &AiOpenai{
		client: openai.NewClient(
			option.WithBaseURL(config.BaseURL),
			option.WithAPIKey(config.ApiKey),
		),
		model: config.Model,
	}
}

const systemOpenAISummaryContent = "我是博客博主 我会给你一篇文章 以博主为主语 返回我文章摘要即可 字数在200左右 使用中文回复"

func (a *AiOpenai) GetSummary(articleContent string) (string, error) {
	chatCompletion, err := a.client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(articleContent),
			openai.SystemMessage(systemOpenAISummaryContent),
		}),
		Model: openai.F(a.model),
	})
	if err != nil {
		return "", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}

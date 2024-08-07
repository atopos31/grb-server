package service

import (
	"context"
	"gvb/config"

	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
)

type AiQianfan struct {
	chat *qianfan.ChatCompletion
}

const systemQianfanStr = "function"

const systemQianfanSummaryContent = "我是博客博主 我会给你一篇文章 以博主为主语 只需返回我文章摘要即可 字数在200左右"

func NewAiQianfan(config config.AiQianfan) *AiQianfan {
	qianfan.GetConfig().AK = config.AccessKey
	qianfan.GetConfig().SK = config.SecretKey

	chat := qianfan.NewChatCompletion(
		qianfan.WithModel(config.Model),
	)
	return &AiQianfan{chat: chat}
}

func (a *AiQianfan) GetSummary(articleContent string) (string, error) {
	resp, err := a.chat.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage(articleContent),
			},
			System: systemQianfanSummaryContent,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Result, nil
}

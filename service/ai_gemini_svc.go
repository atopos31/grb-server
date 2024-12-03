package service

import (
	"context"
	"fmt"
	"gvb/config"
	"log"
	"net/http"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"golang.org/x/net/proxy"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/option"
)

type AIGemini struct {
	client *genai.Client
	model  string
	ctx    context.Context
}

const systemGeminiSummaryContent = "我是博客博主 我会给你一篇文章 以我为主语 返回我文章摘要即可 字数在200左右"

func NewAiGemini(config config.AiGemini) *AIGemini {
	ctx := context.Background()
	dialer, err := proxy.SOCKS5("tcp", config.Sokcs5, nil, proxy.Direct)
	if err != nil {
		panic(err)
	}
	client, err := genai.NewClient(ctx, option.WithHTTPClient(
		&http.Client{
			Transport: &transport.APIKey{
				Key: config.ApiKey,
				Transport: &http.Transport{
					Dial: dialer.Dial,
				},
			},
		},
	),option.WithAPIKey(config.ApiKey))

	return &AIGemini{
		client: client,
		model:  config.Model,
		ctx:    ctx,
	}
}

func (a *AIGemini) GetSummary(articleContent string) (string, error) {
	model := a.client.GenerativeModel(a.model)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemGeminiSummaryContent),
		},
	}
	session := model.StartChat()

	resp, err := session.SendMessage(a.ctx, genai.Text(articleContent))
	if err != nil {
		log.Fatalf("Error sending message: %+v", err)
	}
	var restr strings.Builder
	for _, part := range resp.Candidates[0].Content.Parts {
		restr.WriteString(fmt.Sprintf("%v", part))
	}

	return restr.String(), nil
}

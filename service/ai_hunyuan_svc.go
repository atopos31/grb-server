package service

import (
	"errors"
	"gvb/config"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	hunyuan "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hunyuan/v20230901"
)

type AiHunyuan struct {
	Client  *hunyuan.Client
	Request *hunyuan.ChatCompletionsRequest
}

const endpointURL = "hunyuan.tencentcloudapi.com"

const systemStr = "system"
const userStr = "user"

const systemSummaryContent = "我是博客博主 我会给你一篇文章 以博主为主语 返回我文章摘要即可 字数在300左右"

func NewAiHunyuan(config config.AiHunyuan) *AiHunyuan {
	credential := common.NewCredential(
		config.SecretId,
		config.SecretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = endpointURL

	client, err := hunyuan.NewClient(credential, "", cpf)
	if err != nil {
		panic(err)
	}

	request := hunyuan.NewChatCompletionsRequest()
	request.Model = common.StringPtr(config.Model)

	return &AiHunyuan{
		Client:  client,
		Request: request,
	}
}
func (a *AiHunyuan) GetSummary(articleContent string) (string, error) {
	a.Request.Messages = []*hunyuan.Message{
		{
			Role:    common.StringPtr(systemStr),
			Content: common.StringPtr(systemSummaryContent),
		},
		{
			Role:    common.StringPtr(userStr),
			Content: common.StringPtr(articleContent),
		},
	}
	response, err := a.Client.ChatCompletions(a.Request)
	if err != nil {
		return "", err
	}
	if response.Response != nil {
		return *response.Response.Choices[0].Message.Content, nil
	} else {
		return "", errors.New("response is nil")
	}
}

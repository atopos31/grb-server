package service

import (
	"gvb/config"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	hunyuan "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hunyuan/v20230901"
)

type AiHunyuan struct {
	Credential     *common.Credential
	SummaryRequest *hunyuan.ChatCompletionsRequest
}

func NewAiHunyuan(config config.AiHunyuan) *AiHunyuan {
	summaryRequest := hunyuan.NewChatCompletionsRequest()
	
	return &AiHunyuan{
		Credential: common.NewCredential(
			config.SecretId,
			config.SecretKey,
		),
		SummaryRequest: summaryRequest,
	}
}

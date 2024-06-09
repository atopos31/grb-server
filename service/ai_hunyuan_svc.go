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

// 这里使用值接收者 避免修改SummaryRequest中的内容
func (a AiHunyuan) GetSummary(content string) (string, error) {
	// TODO 对接腾讯混元大模型实现ai生成文章简介
	return "test", nil
}

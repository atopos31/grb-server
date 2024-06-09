package inter

type AiService interface {
	GetSummary(content string) (string, error) // ai获取文章摘要
}

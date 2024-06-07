package inter

type AiService interface {
	GetSummary(content string) (string, error)
}

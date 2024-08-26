package res

type SiteInfo struct {
	ArticleCount  int64 `json:"articlecount"`
	CategoryCount int64 `json:"categorycount"`
	TagCount      int64 `json:"tagcount"`
	ViewsCount    int64 `json:"viewscount"`
}

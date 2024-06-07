package res

type SiteInfo struct {
	ArticleCount  int64  `json:"文章数"`
	CategoryCount int64  `json:"分类数"`
	TagCount      int64  `json:"标签数"`
	ViewsCount    int64  `json:"访问量"`
	Record        string `json:"备案号"`
}

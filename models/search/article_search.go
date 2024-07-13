package search

type ArticleSearch struct {
	Uuid    uint32 `json:"uuid"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

type ArticleSearchResult struct {
	Hits             []ArticleSearch `json:"hits"`
	Total            int64           `json:"total"`
	ProcessingTimeMs int64           `json:"processingTimeMs"`
}

type SearchResponse struct {
	Hits               []Hit       `json:"hits"`
	EstimatedTotalHits int64       `json:"estimatedTotalHits,omitempty"`
	Offset             int64       `json:"offset,omitempty"`
	Limit              int64       `json:"limit,omitempty"`
	ProcessingTimeMs   int64       `json:"processingTimeMs"`
	Query              string      `json:"query"`
	FacetDistribution  interface{} `json:"facetDistribution,omitempty"`
	TotalHits          int64       `json:"totalHits,omitempty"`
	HitsPerPage        int64       `json:"hitsPerPage,omitempty"`
	Page               int64       `json:"page,omitempty"`
	TotalPages         int64       `json:"totalPages,omitempty"`
	FacetStats         interface{} `json:"facetStats,omitempty"`
	IndexUID           string      `json:"indexUid,omitempty"`
}

type Hit struct {
	Formatted Formatted `json:"_formatted"`
	Content   string    `json:"content"`
	Summary   string    `json:"summary"`
	Title     string    `json:"title"`
	UUID      uint32    `json:"uuid"`
}

type Formatted struct {
	Content string `json:"content"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
	UUID    string `json:"uuid"`
}

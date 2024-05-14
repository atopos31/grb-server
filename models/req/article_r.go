package req

type Article struct {
	Title      string   `json:"title" binding:"required"`
	Content    string   `json:"content" binding:"required"`
	Summary    string   `json:"summary"`
	CoverImage string   `json:"cover_image"`
	CategoryID uint     `json:"category_id"`
	Tags       []string `json:"tags"`
	Top        uint8    `json:"top" binding:"oneof=0 1"`
	Status     uint8    `json:"status" binding:"oneof=0 1"`
}

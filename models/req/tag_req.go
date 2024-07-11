package req

type Tag struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}

type TagList struct {
	PageSize int `form:"page_size" binding:"gte=1,lte=20"`
	PageNum  int `form:"page_num" binding:"gte=1"`
}

package req

type Tag struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}

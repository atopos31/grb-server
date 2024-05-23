package req

type Cate struct {
	ID   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}

package res

type Comment struct {
	ID           int            `json:"id"`
	Content      string         `json:"content"`
	UserName     string         `json:"userName"`
	Avatar       string         `json:"avatar"`
	Email        string         `json:"email"`
	CreatedAt    string         `json:"createdAt"`
	ChildComment []ChildComment `json:"childComment"`
}

type ChildComment struct {
	ID        int    `json:"id"`
	ParentID  *int   `json:"parentId"` // 父评论ID 为空为二级评论 不为空为二级评论的回复
	Content   string `json:"content"`
	UserName  string `json:"userName"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

package res

type CommentTree struct {
	CommentId int            `json:"cid"`
	Content   string         `json:"content"`
	UserName  string         `json:"userName"`
	Avatar    string         `json:"avatar"`
	Email     string         `json:"email"`
	CreatedAt string         `json:"createdAt"`
	Children  []*CommentTree `json:"children"`
}

package req

type Comment struct {
	ArticleID uint32 `json:"article_id" binding:"required"`
	Content   string `json:"content" binding:"required"`
	ParentID  *uint  `json:"parent_id"`
	RootID    *uint  `json:"root_id"`
	UserName  string `json:"user_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

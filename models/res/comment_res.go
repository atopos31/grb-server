package res

import "gorm.io/gorm"

type Comment struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	WebSite  string `json:"web_site"`

	CreatedAt    LocalTime      `json:"createdAt"`
	ChildComment []ChildComment `gorm:"-" json:"child_comment"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"` //用来避免查询到软删除的数据
}

type ChildComment struct {
	ID        int            `json:"id"`
	ParentID  *int           `json:"parentId"` // 父评论ID 为空为二级评论 不为空为二级评论的回复
	Content   string         `json:"content"`
	UserName  string         `json:"userName"`
	Avatar    string         `json:"avatar"`
	Email     string         `json:"email"`
	WebSite   string         `json:"web_site"`
	CreatedAt LocalTime      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` //用来避免查询到软删除的数据
}

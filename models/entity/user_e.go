package entity

import "gorm.io/gorm"

type AdminUser struct {
	gorm.Model        // 模型
	Name       string `gorm:"type:varchar(255);not null;default:'admin'"` // 用户名
	Email      string `gorm:"type:varchar(255);not null"`                 // 邮箱
	Password   string `gorm:"type:varchar(255);not null"`                 // 密码
	Avatar     string `gorm:"type:varchar(255);not null"`                 // 头像
	Signature  string `gorm:"type:varchar(255)"`                          // 签名
}

// 设置表名
func (AdminUser) TableName() string {
	return "admin_user"
}

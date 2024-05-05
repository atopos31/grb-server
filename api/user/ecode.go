package user

import "gvb/models/errcode"

var (
	ErrPasswordIncorrect = errcode.NewError(20102, "账号或密码错误")
)

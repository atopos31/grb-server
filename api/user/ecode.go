package user

import "gvb/models/errcode"

var ErrPasswordIncorrect = errcode.NewError(20001, "账号或密码错误")

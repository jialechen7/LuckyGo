package model

import (
	"github.com/jialechen7/go-lottery/common/xerr"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrUserAlreadyRegisterError = xerr.NewErrMsg("用户已注册")
var ErrUserNotExistsError = xerr.NewErrMsg("用户不存在")
var ErrPasswordNotMatch = xerr.NewErrMsg("密码不正确")
var ErrLogin = xerr.NewErrMsg("登录失败")

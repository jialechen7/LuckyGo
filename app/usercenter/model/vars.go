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
var ErrWxMiniAuth = xerr.NewErrMsg("微信登录失败")
var ErrParingWxMiniData = xerr.NewErrMsg("解析微信小程序数据失败")
var ErrGetUserAuth = xerr.NewErrMsg("获取用户授权信息失败")
var ErrGetUserInfo = xerr.NewErrMsg("获取用户信息失败")
var ErrGenerateToken = xerr.NewErrMsg("生成token失败")
var ErrUpdateUserInfo = xerr.NewErrMsg("更新用户信息失败")
var ErrAddUserSponsor = xerr.NewErrMsg("添加赞助商失败")
var ErrDeleteUserSponsor = xerr.NewErrMsg("删除赞助商失败")
var ErrGetUserSponsorDetail = xerr.NewErrMsg("获取赞助商详情失败")
var ErrUpdateUserSponsor = xerr.NewErrMsg("更新赞助商失败")
var ErrGetLotteryStatistic = xerr.NewErrMsg("获取抽奖统计失败")

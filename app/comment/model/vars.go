package model

import (
	"github.com/jialechen7/go-lottery/common/xerr"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrInsertComment = xerr.NewErrMsg("创建评论失败")
var ErrInsertPraise = xerr.NewErrMsg("创建点赞失败")
var ErrDeleteComment = xerr.NewErrMsg("删除评论失败")
var ErrDeletePraise = xerr.NewErrMsg("删除点赞失败")
var ErrUpdateComment = xerr.NewErrMsg("更新评论失败")
var ErrUpdatePraise = xerr.NewErrMsg("更新点赞失败")
var ErrGetCommentDetail = xerr.NewErrMsg("获取评论详情失败")
var ErrGetLastCommentId = xerr.NewErrMsg("获取最后一条评论id失败")
var ErrGetCommentList = xerr.NewErrMsg("获取评论列表失败")
var ErrGetUserInfos = xerr.NewErrMsg("获取用户信息失败")

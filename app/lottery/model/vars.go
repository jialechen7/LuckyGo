package model

import (
	"github.com/jialechen7/go-lottery/common/xerr"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrRowsAffectedZero = xerr.NewErrMsg("操作数据行数为0")
var ErrSearchList = xerr.NewErrMsg("查询列表失败")
var ErrCreateLottery = xerr.NewErrMsg("创建抽奖失败")
var ErrCreatePrize = xerr.NewErrMsg("创建奖品失败")
var ErrLotteryDetail = xerr.NewErrMsg("查询抽奖详情失败")
var ErrFindPrizes = xerr.NewErrMsg("查询奖品失败")
var ErrGetLotteryWinList = xerr.NewErrMsg("查询中奖列表失败")
var ErrGetUserAllLottery = xerr.NewErrMsg("查询用户所有抽奖失败")
var ErrGetUserCreatedLottery = xerr.NewErrMsg("查询用户创建的抽奖失败")
var ErrGetLotteryParticipation = xerr.NewErrMsg("查询抽奖参与者失败")
var ErrGetUserInfo = xerr.NewErrMsg("查询用户信息失败")
var ErrCheckIsParticipated = xerr.NewErrMsg("查询是否参与抽奖失败")
var ErrParticipateLottery = xerr.NewErrMsg("参与抽奖失败")
var ErrHasAnnounced = xerr.NewErrMsg("抽奖已公布，不能抽奖")

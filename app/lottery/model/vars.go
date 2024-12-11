package model

import (
	"github.com/jialechen7/go-lottery/common/xerr"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrSearchList = xerr.NewErrMsg("查询列表失败")
var ErrCreateLottery = xerr.NewErrMsg("创建抽奖失败")
var ErrCreatePrize = xerr.NewErrMsg("创建奖品失败")
var ErrLotteryDetail = xerr.NewErrMsg("查询抽奖详情失败")
var ErrFindPrizes = xerr.NewErrMsg("查询奖品失败")

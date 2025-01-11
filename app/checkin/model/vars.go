package model

import (
	"github.com/jialechen7/go-lottery/common/xerr"
	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrCheckin = xerr.NewErrMsg("签到失败")
var ErrClaimReward = xerr.NewErrMsg("领取奖励失败")
var ErrUpdateSub = xerr.NewErrMsg("更新订阅状态失败")

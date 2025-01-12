package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/jobtype"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/usercenter"
	"github.com/jialechen7/go-lottery/common/wxnotice"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"strconv"
	"time"

	"github.com/jialechen7/go-lottery/app/notice/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrNotifyWishCheckin = xerr.NewErrMsg("通知心愿签到失败")

type NoticeWishSignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeWishSignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeWishSignLogic {
	return &NoticeWishSignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NoticeWishSignLogic) NoticeWishSign(in *pb.NoticeWishSignInReq) (*pb.NoticeWishSignInResp, error) {
	pbUserAuth, err := l.svcCtx.UsercenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserId{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyWishCheckin, "NoticeWishSign GetUserAuthByUserId err:%+v, userId:%d", err, in.UserId)
	}

	if pbUserAuth.UserAuth.AuthKey == "" || pbUserAuth.UserAuth == nil {
		logx.WithContext(l.ctx).Errorw("NoticeWishSign user has no wechat auth",
			logx.Field("userId", in.UserId))
		return nil, errors.Wrapf(ErrNotifyWishCheckin, "NoticeWishSign user has no wechat auth, userId:%d", in.UserId)
	}

	pageAddr := fmt.Sprintf("pages/index/checkin")

	msg := wxnotice.MessageWishCheckin{
		ActivityName:          wxnotice.Item{Value: "心愿签到"},
		ContinuousCheckinDays: wxnotice.Item{Value: strconv.Itoa(int(in.Accumulate))},
		Time:                  wxnotice.Item{Value: time.Now().Format(time.DateOnly)},
		CheckinAward:          wxnotice.Item{Value: fmt.Sprintf("今日签到可获得%s心愿值", strconv.Itoa(int(in.Reward)))},
		RemindText:            wxnotice.Item{Value: "每日签到赢取心愿值"},
	}

	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyWishCheckin, "NoticeWishSign json.Marshal err:%+v", err)
	}

	payload := jobtype.WxMiniProgramNotifyUserPayload{
		MsgType:  msg.Type(),
		OpenId:   pbUserAuth.UserAuth.AuthKey,
		PageAddr: pageAddr,
		Data:     string(jsonBytes),
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyWishCheckin, "NoticeWishSign payload json.Marshal err:%+v", err)
	}
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgWxMiniProgramNotifyUser, payloadBytes))
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyWishCheckin, "NoticeWishSign AsynqClient.Enqueue err:%+v", err)
	}
	return &pb.NoticeWishSignInResp{}, nil
}

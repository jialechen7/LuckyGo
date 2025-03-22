package logic

import (
	"context"
	"encoding/json"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/subscribeMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/hibiken/asynq"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/internal/svc"
	"github.com/jialechen7/go-lottery/app/mqueue/cmd/job/jobtype"
	"github.com/jialechen7/go-lottery/common/wxnotice"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

// 微信返回的错误码
const WxErrCodeUserRefuseReceiveMsg = 43101

var ErrNotifyUserFail = xerr.NewErrMsg("通知失败")

// WxMiniProgramNotifyUserHandler mini program notify user
type WxMiniProgramNotifyUserHandler struct {
	svcCtx *svc.ServiceContext
}

func NewWxMiniProgramNotifyUserHandler(svcCtx *svc.ServiceContext) *WxMiniProgramNotifyUserHandler {
	return &WxMiniProgramNotifyUserHandler{
		svcCtx: svcCtx,
	}
}

func (l *WxMiniProgramNotifyUserHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var (
		err        error
		p          jobtype.WxMiniProgramNotifyUserPayload
		templateId string
	)
	if err = json.Unmarshal(t.Payload(), &p); err != nil {
		// 不可重试的错误，记录日志
		logx.Error("WxMiniProgramNotifyUserHandler ProcessTask payload json unmarshal err",
			logx.Field("err", err))
		return nil
	}
	data := new(power.HashMap)

	switch p.MsgType {
	case wxnotice.TypeWishSign:
		var msg wxnotice.MessageWishCheckin
		err := json.Unmarshal([]byte(p.Data), &msg)
		if err != nil {
			// 不可重试的错误，记录日志
			logx.Error("WxMiniProgramNotifyUserHandler ProcessTask data json unmarshal err",
				logx.Field("err", err),
				logx.Field("data", p.Data))
			return nil
		}
		templateId = msg.TemplateId()
		data, err = power.StructToHashMap(&msg)
		if err != nil {
			// 不可重试的错误，记录日志
			logx.Error("WxMiniProgramNotifyUserHandler ProcessTask data convert err",
				logx.Field("err", err),
				logx.Field("data", data))
			return nil
		}
	default:
		logx.Error("WxMiniProgramNotifyUserHandler ProcessTask payload data invalid",
			logx.Field("payload", p))
	}

	reqData := &request.RequestSubscribeMessageSend{
		ToUser:           p.OpenId,
		TemplateID:       templateId,
		Page:             p.PageAddr,
		MiniProgramState: "developer",
		Lang:             "zh_CN",
		Data:             data,
	}
	// 发送消息
	resp, err := l.svcCtx.WxMiniProgram.SubscribeMessage.Send(ctx, reqData)

	if err != nil {
		// 可重试的错误
		return errors.Wrapf(ErrNotifyUserFail, "WxMiniProgramNotifyUserHandler send message err:%v, reqData:%+v", err, reqData)
	}

	if resp.ErrCode != 0 {
		switch resp.ErrCode {
		// 不可重试的错误码
		case WxErrCodeUserRefuseReceiveMsg:
			logx.Infow("WxMiniProgramNotifyUserHandler user refuse receive msg",
				logx.Field("openid", p.OpenId),
			)
			return nil
		default:
			// 可重试的错误码，后续进行细分
			return errors.Wrapf(ErrNotifyUserFail, "WxMiniProgramNotifyUserHandler send message fail, errCode:%v, errMsg: %v, reqData:%+v", resp.ErrCode, resp.ErrMsg, reqData)
		}
	}

	return nil
}

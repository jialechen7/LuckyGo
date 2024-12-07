package user

import (
	"context"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/api/internal/types"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/app/usercenter/model"
	"github.com/jialechen7/go-lottery/common/constants"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 微信登录注册
func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMiniAuthReq) (resp *types.WXMiniAuthResp, err error) {
	// Wechat-Mini
	miniProgram := wechat.NewWechat().GetMiniProgram(&config.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.AppId,
		AppSecret: l.svcCtx.Config.WxMiniConf.AppSecret,
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniProgram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		return nil, errors.Wrapf(model.ErrWxMiniAuth, "Wechat-Mini auth fail, err: %v, code: %s, authResult: %+v", err, req.Code, authResult)
	}

	// Parsing WeChat-Mini return data
	userData, err := miniProgram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Wrapf(model.ErrParingWxMiniData, "Parsing WeChat-Mini data fail, req: %+v, err: %v, authResult: %+v", req, err, authResult)
	}
	pbUserAuth, err := l.svcCtx.UsercenterRpc.GetUserAuthByAuthKey(l.ctx, &pb.GetUserAuthByAuthKeyReq{
		AuthKey:  userData.OpenID,
		AuthType: constants.UserAuthTypeWxMini,
	})
	if err != nil {
		return nil, errors.Wrapf(model.ErrGetUserAuth, "Get user auth fail, err: %v, req: %+v", err, req)
	}

	resp = &types.WXMiniAuthResp{}
	if pbUserAuth == nil || pbUserAuth.UserAuth.Id == 0 {
		if len(req.Nickname) == 0 {
			nicknameArr := []string{userData.NickName, utility.Krand(6, utility.KC_RAND_KIND_NUM)}
			nickName := strings.Join(nicknameArr, "#")
			req.Nickname = nickName
		}
		if len(req.Avatar) == 0 {
			req.Avatar = userData.AvatarURL
		}

		pbWxMiniAuthResp, err := l.svcCtx.UsercenterRpc.WxMiniAuth(l.ctx, &pb.WxMiniAuthReq{
			Nickname: req.Nickname,
			Avatar:   req.Avatar,
			AuthType: constants.UserAuthTypeWxMini,
			AuthKey:  userData.OpenID,
		})
		if err != nil {
			return nil, errors.Wrapf(model.ErrWxMiniAuth, "Wechat-Mini auth fail, err: %v, req: %+v", err, req)
		}

		_ = copier.Copy(resp, pbWxMiniAuthResp)
	} else {
		userId := pbUserAuth.UserAuth.UserId
		pbToken, err := l.svcCtx.UsercenterRpc.GenerateToken(l.ctx, &pb.GenerateTokenReq{
			UserId: userId,
		})
		if err != nil {
			return nil, err
		}

		_ = copier.Copy(resp, pbToken)
	}
	return resp, nil
}

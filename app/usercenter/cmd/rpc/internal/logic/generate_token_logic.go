package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessToken, err :=
		utility.GenerateJWT(l.svcCtx.Config.JwtAuth.AccessSecret, now, l.svcCtx.Config.JwtAuth.AccessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "GenerateToken err:%v", err)
	}
	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: l.svcCtx.Config.JwtAuth.AccessExpire,
		RefreshAfter: l.svcCtx.Config.JwtAuth.AccessExpire / 2,
	}, nil
}

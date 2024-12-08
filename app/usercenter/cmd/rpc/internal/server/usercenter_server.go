// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: usercenter.proto

package server

import (
	"context"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/logic"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UsercenterServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UsercenterServer) WxMiniAuth(ctx context.Context, in *pb.WxMiniAuthReq) (*pb.WxMiniAuthResp, error) {
	l := logic.NewWxMiniAuthLogic(ctx, s.svcCtx)
	return l.WxMiniAuth(in)
}

func (s *UsercenterServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

func (s *UsercenterServer) GetUserAuthByAuthKey(ctx context.Context, in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {
	l := logic.NewGetUserAuthByAuthKeyLogic(ctx, s.svcCtx)
	return l.GetUserAuthByAuthKey(in)
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UsercenterServer) UpdateUserBaseInfo(ctx context.Context, in *pb.UpdateUserBaseInfoReq) (*pb.UpdateUserBaseInfoResp, error) {
	l := logic.NewUpdateUserBaseInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserBaseInfo(in)
}

// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: usercenter.proto

package usercenter

import (
	"context"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddUserReq               = pb.AddUserReq
	AddUserResp              = pb.AddUserResp
	AddUserSponsorReq        = pb.AddUserSponsorReq
	AddUserSponsorResp       = pb.AddUserSponsorResp
	DelUserSponsorReq        = pb.DelUserSponsorReq
	DelUserSponsorResp       = pb.DelUserSponsorResp
	GenerateTokenReq         = pb.GenerateTokenReq
	GenerateTokenResp        = pb.GenerateTokenResp
	GetUserAuthByAuthKeyReq  = pb.GetUserAuthByAuthKeyReq
	GetUserAuthByAuthKeyResp = pb.GetUserAuthByAuthKeyResp
	GetUserAuthByUserId      = pb.GetUserAuthByUserId
	GetUserAuthByUserIdResp  = pb.GetUserAuthByUserIdResp
	GetUserInfoByUserIdsReq  = pb.GetUserInfoByUserIdsReq
	GetUserInfoByUserIdsResp = pb.GetUserInfoByUserIdsResp
	GetUserInfoReq           = pb.GetUserInfoReq
	GetUserInfoResp          = pb.GetUserInfoResp
	LoginReq                 = pb.LoginReq
	LoginResp                = pb.LoginResp
	RegisterReq              = pb.RegisterReq
	RegisterResp             = pb.RegisterResp
	SearchUserSponsorReq     = pb.SearchUserSponsorReq
	SearchUserSponsorResp    = pb.SearchUserSponsorResp
	SponsorDetailReq         = pb.SponsorDetailReq
	SponsorDetailResp        = pb.SponsorDetailResp
	UpdateUserBaseInfoReq    = pb.UpdateUserBaseInfoReq
	UpdateUserBaseInfoResp   = pb.UpdateUserBaseInfoResp
	UpdateUserSponsorReq     = pb.UpdateUserSponsorReq
	UpdateUserSponsorResp    = pb.UpdateUserSponsorResp
	User                     = pb.User
	UserAuth                 = pb.UserAuth
	UserInfoForComment       = pb.UserInfoForComment
	UserSponsor              = pb.UserSponsor
	WxMiniAuthReq            = pb.WxMiniAuthReq
	WxMiniAuthResp           = pb.WxMiniAuthResp

	Usercenter interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		WxMiniAuth(ctx context.Context, in *WxMiniAuthReq, opts ...grpc.CallOption) (*WxMiniAuthResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		UpdateUserBaseInfo(ctx context.Context, in *UpdateUserBaseInfoReq, opts ...grpc.CallOption) (*UpdateUserBaseInfoResp, error)
		GetUserInfoByUserIds(ctx context.Context, in *GetUserInfoByUserIdsReq, opts ...grpc.CallOption) (*GetUserInfoByUserIdsResp, error)
		GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error)
		GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserId, opts ...grpc.CallOption) (*GetUserAuthByUserIdResp, error)
		AddUserSponsor(ctx context.Context, in *AddUserSponsorReq, opts ...grpc.CallOption) (*AddUserSponsorResp, error)
		UpdateUserSponsor(ctx context.Context, in *UpdateUserSponsorReq, opts ...grpc.CallOption) (*UpdateUserSponsorResp, error)
		DelUserSponsor(ctx context.Context, in *DelUserSponsorReq, opts ...grpc.CallOption) (*DelUserSponsorResp, error)
		SearchUserSponsor(ctx context.Context, in *SearchUserSponsorReq, opts ...grpc.CallOption) (*SearchUserSponsorResp, error)
		SponsorDetail(ctx context.Context, in *SponsorDetailReq, opts ...grpc.CallOption) (*SponsorDetailResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsercenter) WxMiniAuth(ctx context.Context, in *WxMiniAuthReq, opts ...grpc.CallOption) (*WxMiniAuthResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.WxMiniAuth(ctx, in, opts...)
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserBaseInfo(ctx context.Context, in *UpdateUserBaseInfoReq, opts ...grpc.CallOption) (*UpdateUserBaseInfoResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserBaseInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserInfoByUserIds(ctx context.Context, in *GetUserInfoByUserIdsReq, opts ...grpc.CallOption) (*GetUserInfoByUserIdsResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfoByUserIds(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthByAuthKey(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserId, opts ...grpc.CallOption) (*GetUserAuthByUserIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthByUserId(ctx, in, opts...)
}

func (m *defaultUsercenter) AddUserSponsor(ctx context.Context, in *AddUserSponsorReq, opts ...grpc.CallOption) (*AddUserSponsorResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.AddUserSponsor(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateUserSponsor(ctx context.Context, in *UpdateUserSponsorReq, opts ...grpc.CallOption) (*UpdateUserSponsorResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateUserSponsor(ctx, in, opts...)
}

func (m *defaultUsercenter) DelUserSponsor(ctx context.Context, in *DelUserSponsorReq, opts ...grpc.CallOption) (*DelUserSponsorResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.DelUserSponsor(ctx, in, opts...)
}

func (m *defaultUsercenter) SearchUserSponsor(ctx context.Context, in *SearchUserSponsorReq, opts ...grpc.CallOption) (*SearchUserSponsorResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SearchUserSponsor(ctx, in, opts...)
}

func (m *defaultUsercenter) SponsorDetail(ctx context.Context, in *SponsorDetailReq, opts ...grpc.CallOption) (*SponsorDetailResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SponsorDetail(ctx, in, opts...)
}

// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: lottery.proto

package lottery

import (
	"context"

	"github.com/jialechen7/go-lottery/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddLotteryParticipationReq             = pb.AddLotteryParticipationReq
	AddLotteryParticipationResp            = pb.AddLotteryParticipationResp
	AddLotteryReq                          = pb.AddLotteryReq
	AddLotteryResp                         = pb.AddLotteryResp
	AnnounceLotteryReq                     = pb.AnnounceLotteryReq
	AnnounceLotteryResp                    = pb.AnnounceLotteryResp
	CheckIsParticipatedReq                 = pb.CheckIsParticipatedReq
	CheckIsParticipatedResp                = pb.CheckIsParticipatedResp
	CheckLotteryCreatedReq                 = pb.CheckLotteryCreatedReq
	CheckLotteryCreatedResp                = pb.CheckLotteryCreatedResp
	CheckLotteryParticipatedReq            = pb.CheckLotteryParticipatedReq
	CheckLotteryParticipatedResp           = pb.CheckLotteryParticipatedResp
	CheckUserIsWonReq                      = pb.CheckUserIsWonReq
	CheckUserIsWonResp                     = pb.CheckUserIsWonResp
	ClockTask                              = pb.ClockTask
	GetLotteryListAfterLoginReq            = pb.GetLotteryListAfterLoginReq
	GetLotteryListAfterLoginResp           = pb.GetLotteryListAfterLoginResp
	GetLotteryStatisticReq                 = pb.GetLotteryStatisticReq
	GetLotteryStatisticResp                = pb.GetLotteryStatisticResp
	GetParticipationUserIdsByLotteryIdReq  = pb.GetParticipationUserIdsByLotteryIdReq
	GetParticipationUserIdsByLotteryIdResp = pb.GetParticipationUserIdsByLotteryIdResp
	GetPrizeListByLotteryIdReq             = pb.GetPrizeListByLotteryIdReq
	GetPrizeListByLotteryIdResp            = pb.GetPrizeListByLotteryIdResp
	GetUserAllListReq                      = pb.GetUserAllListReq
	GetUserAllListResp                     = pb.GetUserAllListResp
	GetUserCreatedListReq                  = pb.GetUserCreatedListReq
	GetUserCreatedListResp                 = pb.GetUserCreatedListResp
	GetUserWonListReq                      = pb.GetUserWonListReq
	GetUserWonListResp                     = pb.GetUserWonListResp
	GetWonListByLotteryIdReq               = pb.GetWonListByLotteryIdReq
	GetWonListByLotteryIdResp              = pb.GetWonListByLotteryIdResp
	GetWonListCountReq                     = pb.GetWonListCountReq
	GetWonListCountResp                    = pb.GetWonListCountResp
	Lottery                                = pb.Lottery
	LotteryDetailReq                       = pb.LotteryDetailReq
	LotteryDetailResp                      = pb.LotteryDetailResp
	LotteryParticipation                   = pb.LotteryParticipation
	Prize                                  = pb.Prize
	SearchLotteryParticipationReq          = pb.SearchLotteryParticipationReq
	SearchLotteryParticipationResp         = pb.SearchLotteryParticipationResp
	SearchLotteryReq                       = pb.SearchLotteryReq
	SearchLotteryResp                      = pb.SearchLotteryResp
	UserInfo                               = pb.UserInfo
	UserLotteryList                        = pb.UserLotteryList
	WonList                                = pb.WonList

	LotteryZrpcClient interface {
		SearchLottery(ctx context.Context, in *SearchLotteryReq, opts ...grpc.CallOption) (*SearchLotteryResp, error)
		GetLotteryListAfterLogin(ctx context.Context, in *GetLotteryListAfterLoginReq, opts ...grpc.CallOption) (*GetLotteryListAfterLoginResp, error)
		AddLottery(ctx context.Context, in *AddLotteryReq, opts ...grpc.CallOption) (*AddLotteryResp, error)
		LotteryDetail(ctx context.Context, in *LotteryDetailReq, opts ...grpc.CallOption) (*LotteryDetailResp, error)
		AnnounceLottery(ctx context.Context, in *AnnounceLotteryReq, opts ...grpc.CallOption) (*AnnounceLotteryResp, error)
		GetUserWonList(ctx context.Context, in *GetUserWonListReq, opts ...grpc.CallOption) (*GetUserWonListResp, error)
		GetWonListByLotteryId(ctx context.Context, in *GetWonListByLotteryIdReq, opts ...grpc.CallOption) (*GetWonListByLotteryIdResp, error)
		GetUserAllList(ctx context.Context, in *GetUserAllListReq, opts ...grpc.CallOption) (*GetUserAllListResp, error)
		GetUserCreatedList(ctx context.Context, in *GetUserCreatedListReq, opts ...grpc.CallOption) (*GetUserCreatedListResp, error)
		SearchLotteryParticipation(ctx context.Context, in *SearchLotteryParticipationReq, opts ...grpc.CallOption) (*SearchLotteryParticipationResp, error)
		GetLotteryStatistic(ctx context.Context, in *GetLotteryStatisticReq, opts ...grpc.CallOption) (*GetLotteryStatisticResp, error)
		AddLotteryParticipation(ctx context.Context, in *AddLotteryParticipationReq, opts ...grpc.CallOption) (*AddLotteryParticipationResp, error)
		CheckLotteryParticipated(ctx context.Context, in *CheckLotteryParticipatedReq, opts ...grpc.CallOption) (*CheckLotteryParticipatedResp, error)
		CheckLotteryCreated(ctx context.Context, in *CheckLotteryCreatedReq, opts ...grpc.CallOption) (*CheckLotteryCreatedResp, error)
	}

	defaultLotteryZrpcClient struct {
		cli zrpc.Client
	}
)

func NewLotteryZrpcClient(cli zrpc.Client) LotteryZrpcClient {
	return &defaultLotteryZrpcClient{
		cli: cli,
	}
}

func (m *defaultLotteryZrpcClient) SearchLottery(ctx context.Context, in *SearchLotteryReq, opts ...grpc.CallOption) (*SearchLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.SearchLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetLotteryListAfterLogin(ctx context.Context, in *GetLotteryListAfterLoginReq, opts ...grpc.CallOption) (*GetLotteryListAfterLoginResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetLotteryListAfterLogin(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) AddLottery(ctx context.Context, in *AddLotteryReq, opts ...grpc.CallOption) (*AddLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.AddLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) LotteryDetail(ctx context.Context, in *LotteryDetailReq, opts ...grpc.CallOption) (*LotteryDetailResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.LotteryDetail(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) AnnounceLottery(ctx context.Context, in *AnnounceLotteryReq, opts ...grpc.CallOption) (*AnnounceLotteryResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.AnnounceLottery(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetUserWonList(ctx context.Context, in *GetUserWonListReq, opts ...grpc.CallOption) (*GetUserWonListResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetUserWonList(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetWonListByLotteryId(ctx context.Context, in *GetWonListByLotteryIdReq, opts ...grpc.CallOption) (*GetWonListByLotteryIdResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetWonListByLotteryId(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetUserAllList(ctx context.Context, in *GetUserAllListReq, opts ...grpc.CallOption) (*GetUserAllListResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetUserAllList(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetUserCreatedList(ctx context.Context, in *GetUserCreatedListReq, opts ...grpc.CallOption) (*GetUserCreatedListResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetUserCreatedList(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) SearchLotteryParticipation(ctx context.Context, in *SearchLotteryParticipationReq, opts ...grpc.CallOption) (*SearchLotteryParticipationResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.SearchLotteryParticipation(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) GetLotteryStatistic(ctx context.Context, in *GetLotteryStatisticReq, opts ...grpc.CallOption) (*GetLotteryStatisticResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.GetLotteryStatistic(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) AddLotteryParticipation(ctx context.Context, in *AddLotteryParticipationReq, opts ...grpc.CallOption) (*AddLotteryParticipationResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.AddLotteryParticipation(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) CheckLotteryParticipated(ctx context.Context, in *CheckLotteryParticipatedReq, opts ...grpc.CallOption) (*CheckLotteryParticipatedResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.CheckLotteryParticipated(ctx, in, opts...)
}

func (m *defaultLotteryZrpcClient) CheckLotteryCreated(ctx context.Context, in *CheckLotteryCreatedReq, opts ...grpc.CallOption) (*CheckLotteryCreatedResp, error) {
	client := pb.NewLotteryClient(m.cli.Conn())
	return client.CheckLotteryCreated(ctx, in, opts...)
}

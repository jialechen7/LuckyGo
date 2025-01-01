package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPraiseByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPraiseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPraiseByIdLogic {
	return &GetPraiseByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPraiseByIdLogic) GetPraiseById(in *pb.GetPraiseByIdReq) (*pb.GetPraiseByIdResp, error) {
	dbPraise, err := l.svcCtx.PraiseModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_PRAISE_ERROR), "Failed to find praise, PraiseModel FindOne fail , req : %+v , err : %v", in, err)
	}

	pbPraise := &pb.Praise{}
	_ = copier.Copy(pbPraise, dbPraise)

	return &pb.GetPraiseByIdResp{
		Praise: pbPraise,
	}, nil
}

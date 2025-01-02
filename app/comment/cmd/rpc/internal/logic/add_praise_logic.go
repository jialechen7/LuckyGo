package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/comment/model"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPraiseLogic {
	return &AddPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPraiseLogic) AddPraise(in *pb.AddPraiseReq) (*pb.AddPraiseResp, error) {
	dbPraise := &model.Praise{}
	_ = copier.Copy(dbPraise, in)
	err := l.svcCtx.PraiseModel.Insert(l.ctx, nil, dbPraise)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERT_COMMENT_ERROR), "AddPraise error: %v", err)
	}

	return &pb.AddPraiseResp{}, nil
}

package logic

import (
	"context"
	"github.com/jialechen7/go-lottery/app/usercenter/model"
	"github.com/jialechen7/go-lottery/common/utility"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/internal/svc"
	"github.com/jialechen7/go-lottery/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// check if user exists
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "mobile:%s,err:%+v", in.Mobile, err)
	}
	if user != nil {
		return nil, errors.Wrapf(model.ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%+v", in.Mobile, err)
	}

	user = &model.User{}
	// insert user and user_auth
	err = l.svcCtx.TransactCtx(l.ctx, func(db *gorm.DB) error {
		user.Mobile = in.Mobile
		if len(user.Nickname) == 0 {
			user.Nickname = utility.Krand(8, utility.KC_RAND_KIND_ALL)
		}
		if len(in.Password) > 0 {
			user.Password = utility.Md5ByString(in.Password)
		}

		err := l.svcCtx.UserModel.Insert(l.ctx, db, user)
		if err != nil {
			return errors.Wrapf(err, "Register db user Insert err:%v", err)
		}

		userAuth := new(model.UserAuth)
		userAuth.UserId = user.Id
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		err = l.svcCtx.UserAuthModel.Insert(l.ctx, db, userAuth)
		if err != nil {
			return errors.Wrapf(err, "Register db user_auth Insert err:%v", err)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Register db transaction error")
	}

	// generate token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	pbResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Register GenerateToken err userId:%d , err:%v", user.Id, err)
	}
	return &pb.RegisterResp{
		AccessToken:  pbResp.AccessToken,
		AccessExpire: pbResp.AccessExpire,
		RefreshAfter: pbResp.RefreshAfter,
	}, nil
}

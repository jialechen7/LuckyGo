package event

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/svc"
	"github.com/jialechen7/go-lottery/app/notice/cmd/api/internal/types"
	"github.com/jialechen7/go-lottery/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"sort"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrWxEventSignatureInvalidError = xerr.NewErrMsg("wechat event signature is invalid")
)

type VerifyEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证小程序回调消息
func NewVerifyEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEventLogic {
	return &VerifyEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyEventLogic) VerifyEvent(req *types.VerifyEventReq, w http.ResponseWriter) (resp *types.VerifyEventResp, err error) {
	// 参考链接：https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html
	tmpSlice := []string{l.svcCtx.Config.WxMessageNoticeConf.Token, req.Timestamp, req.Nonce}
	sort.Slice(tmpSlice, func(i, j int) bool {
		return tmpSlice[i] < tmpSlice[j]
	})
	str := strings.Join(tmpSlice, "")
	h := sha1.New()
	h.Write([]byte(str))
	sha1Str := hex.EncodeToString(h.Sum(nil))
	if sha1Str != req.Signature {
		return nil, errors.Wrapf(ErrWxEventSignatureInvalidError, "Verify event err : %v ,req:%+v", err, req)
	}

	w.Header().Set(httpx.ContentType, "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(req.Echostr))
	if err != nil {
		return nil, errors.Wrapf(err, "Write event resp error : %v ,req:%+v", err, req)
	}
	return
}

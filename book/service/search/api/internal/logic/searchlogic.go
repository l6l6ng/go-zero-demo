package logic

import (
	"context"

	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/svc"
	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {

	logx.Infof("userId: %v", l.ctx.Value("userId"))
	return &types.SearchReply{
		Name:  "1111",
		Count: 10,
	}, nil
}

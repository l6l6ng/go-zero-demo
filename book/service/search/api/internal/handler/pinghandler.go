package handler

import (
	"net/http"

	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/logic"
	"github.com/l6l6ng/go-zero-demo/book/service/search/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

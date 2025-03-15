package handler

import (
	"net/http"

	"banking-system/transaction-service/rpc/internal/logic"
	"banking-system/transaction-service/rpc/internal/svc"
	"banking-system/transaction-service/rpc/types/transaction/v1alpha1"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTransactionHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req v1alpha1.GetTransactionHistoryRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetTransactionHistoryLogic(r.Context(), svcCtx)
		resp, err := l.GetTransactionHistory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

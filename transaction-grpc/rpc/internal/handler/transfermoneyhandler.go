package handler

import (
	"net/http"

	"banking-system/transaction-service/rpc/internal/logic"
	"banking-system/transaction-service/rpc/internal/svc"
	"banking-system/transaction-service/rpc/types/transaction/v1alpha1"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TransferMoneyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req v1alpha1.TransferMoneyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTransferMoneyLogic(r.Context(), svcCtx)
		_, err := l.TransferMoney(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

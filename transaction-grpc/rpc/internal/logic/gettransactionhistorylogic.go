package logic

import (
	"context"

	"banking-system/transaction-service/rpc/internal/svc"
	"banking-system/transaction-service/rpc/types/transaction/v1alpha1"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransactionHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTransactionHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransactionHistoryLogic {
	return &GetTransactionHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTransactionHistoryLogic) GetTransactionHistory(in *v1alpha1.GetTransactionHistoryRequest) (*v1alpha1.GetTransactionHistoryResponse, error) {
	
	

	return &v1alpha1.GetTransactionHistoryResponse{}, nil
}

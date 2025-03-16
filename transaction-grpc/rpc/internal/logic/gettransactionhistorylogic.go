package logic

import (
	"context"
	"fmt"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	trs, err := l.svcCtx.DB.GetTransactionHistory(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("error retrieving transaction history: %w", err)
	}

	return trs, nil
}

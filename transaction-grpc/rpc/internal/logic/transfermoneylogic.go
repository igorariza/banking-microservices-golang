package logic

import (
	"context"
	"errors"
	"time"

	model "banking-system/transaction-service/model/mongo/nocache"
	"banking-system/transaction-service/pkg/utils"
	"banking-system/transaction-service/rpc/internal/svc"
	"banking-system/transaction-service/rpc/types/transaction/v1alpha1"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransferMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferMoneyLogic {
	return &TransferMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransferMoneyLogic) TransferMoney(in *v1alpha1.TransferMoneyRequest) (*v1alpha1.TransferMoneyResponse, error) {

	if in.Amount < 0 {
		return nil, errors.New("amount cannot be negative")
	}

	trc, err := l.svcCtx.DB.TransferMoney(context.Background(), in)
	if err != nil {
		return nil, err
	}

	go func() {
		utils.PublishTransactionEvent(context.Background(), &model.Transaction{
			FromAccount: in.FromAccount,
			ToAccount:   in.ToAccount,
			Amount:      float64(in.Amount),
			Timestamp:   time.Now().String(),
		})
	}()

	return &v1alpha1.TransferMoneyResponse{
		Id:          trc.Id,
		FromAccount: in.FromAccount,
		ToAccount:   in.ToAccount,
		Amount:      in.Amount,
		Timestamp:   trc.Timestamp,
	}, nil
}

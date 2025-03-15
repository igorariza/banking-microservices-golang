package logic

import (
	"context"
	"errors"

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

	fromAccount, err := l.TransferMoney(in)
	if err != nil {
		return nil, err
	}

	toAccount, err := l.svcCtx.DB.FindOneByName(context.Background(), in.ToAccount)
	if err != nil {
		return nil, err
	}

	if fromAccount.Amount < in.Amount {
		return nil, errors.New("insufficient balance")
	}

	fromAccount.Amount -= in.Amount
	toAccount.Amount += float64(in.Amount)

	_, err = l.svcCtx.DB.TransferMoney(context.Background(), in)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.DB.UpdateAccountBalance(context.Background(), fromAccount.FromAccount)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.DB.UpdateAccountBalance(context.Background(), in.ToAccount)
	if err != nil {
		return nil, err
	}

	return &v1alpha1.TransferMoneyResponse{
		FromAccount: fromAccount.FromAccount,
		ToAccount:   in.ToAccount,
		Amount:      in.Amount,
	}, nil
}

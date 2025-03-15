package logic

import (
	"context"

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
	// (fromAccountID string, toAccountID string, amount float64) (string, error) {
	// 	if amount < 0 {
	// 		return "", errors.New("amount cannot be negative")
	// 	}
	
	// 	fromAccount, err := s.GetAccountBalance(fromAccountID)
	// 	if err != nil {
	// 		return "", err
	// 	}
	
	// 	toAccount, err := s.GetAccountBalance(toAccountID)
	// 	if err != nil {
	// 		return "", err
	// 	}
	
	// 	if fromAccount < amount {
	// 		return "", errors.New("insufficient balance")
	// 	}
	
	// 	fromAccount -= amount
	// 	toAccount += amount
	
	// 	_, err = s.Client.Database(s.Database.Name()).Collection("accounts").UpdateOne(context.Background(), bson.M{"id": fromAccountID},
	// 		bson.M{"$set": bson.M{"balance": fromAccount}})
	// 	if err != nil {
	// 		return "", err
	// 	}
	
	// 	_, err = s.Client.Database(s.Database.Name()).Collection("accounts").UpdateOne(context.Background(), bson.M{"id": toAccountID},
	// 		bson.M{"$set": bson.M{"balance": toAccount}})
	// 	if err != nil {
	// 		return "", err
	// 	}
	
	// 	return uuid.New().String(), nil

	return &v1alpha1.TransferMoneyResponse{}, nil
}

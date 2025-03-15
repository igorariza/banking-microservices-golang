package model

import (
	"context"

	v1alpha1 "banking-system/transaction-service/rpc/types/transaction/v1alpha1"
	"github.com/zeromicro/go-zero/core/stores/mon"
)

type transactionModel interface {
	TransferMoney(ctx context.Context, data *v1alpha1.TransferMoneyRequest) (*v1alpha1.TransferMoneyResponse, error)
	GetTransactionHistory(ctx context.Context, data *v1alpha1.GetTransactionHistoryRequest) (*v1alpha1.GetTransactionHistoryResponse, error)
	UpdateAccountBalance(ctx context.Context, account string) (*Transaction, error)
	FindOneByName(ctx context.Context, name string) (*Transaction, error)
}

type defaultTransactionModel struct {
	conn *mon.Model
}

func newDefaultTransactionModel(conn *mon.Model) *defaultTransactionModel {
	return &defaultTransactionModel{conn: conn}
}

func (m *defaultTransactionModel) TransferMoney(ctx context.Context, data *v1alpha1.TransferMoneyRequest) (*v1alpha1.TransferMoneyResponse, error) {
	

	return nil, nil
}

func (m *defaultTransactionModel) GetTransactionHistory(ctx context.Context, data *v1alpha1.GetTransactionHistoryRequest) (*v1alpha1.GetTransactionHistoryResponse, error) {
	
	return nil, nil
}

func (m *defaultTransactionModel) UpdateAccountBalance(ctx context.Context, account string) (*Transaction, error) {
	//data.UpdateAt = time.Now()

	// _, err := m.conn.ReplaceOne(ctx, bson.M{"_id": data}, data)
	return nil, nil
}


func (m *defaultTransactionModel) FindOneByName(ctx context.Context, name string) (*Transaction, error) {
	var data Transaction

	// err := m.conn.FindOne(ctx, &data, bson.M{"name": name})
	// switch err {
	// case nil:
	// 	return &data, nil
	// case mon.ErrNotFound:
	// 	return nil, ErrNotFound
	// default:
	// 	return nil, err
	// }

	return &data, nil
}

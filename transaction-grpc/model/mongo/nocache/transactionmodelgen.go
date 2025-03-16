package model

import (
	"context"
	"errors"
	"log"
	"time"

	v1alpha1 "banking-system/transaction-service/rpc/types/transaction/v1alpha1"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
)

type transactionModel interface {
	TransferMoney(ctx context.Context, data *v1alpha1.TransferMoneyRequest) (*v1alpha1.TransferMoneyResponse, error)
	GetTransactionHistory(ctx context.Context, data *v1alpha1.GetTransactionHistoryRequest) (*v1alpha1.GetTransactionHistoryResponse, error)
	UpdateAccountBalance(ctx context.Context, account *Transaction) (*Transaction, error)
	FindOneByName(ctx context.Context, name string) (*Transaction, error)
}

type defaultTransactionModel struct {
	conn *mon.Model
}

func newDefaultTransactionModel(conn *mon.Model) *defaultTransactionModel {
	return &defaultTransactionModel{conn: conn}
}

func (m *defaultTransactionModel) TransferMoney(ctx context.Context, data *v1alpha1.TransferMoneyRequest) (*v1alpha1.TransferMoneyResponse, error) {
	fromAccount, err := m.FindOneByName(ctx, data.FromAccount)
	if err != nil {
		return nil, err
	}

	toAccount, err := m.FindOneByName(ctx, data.ToAccount)
	if err != nil {
		return nil, err
	}

	if int64(fromAccount.Amount) < int64(data.Amount) {
		return nil, errors.New("insufficient balance")
	}

	fromAccount.Amount -= float64(data.Amount)
	toAccount.Amount += float64(data.Amount)

	_, err = m.UpdateAccountBalance(ctx, fromAccount)
	if err != nil {
		return nil, err
	}
	_, err = m.UpdateAccountBalance(ctx, toAccount)
	if err != nil {
		return nil, err
	}

	m.conn.Database().Collection("transaction")
	transaction := Transaction{
		ID:          uuid.New().String(),
		FromAccount: data.FromAccount,
		ToAccount:   data.ToAccount,
		Amount:      float64(data.Amount),
		Timestamp:   time.Now().String(),
	}
	_, err = m.conn.InsertOne(ctx, transaction)

	return &v1alpha1.TransferMoneyResponse{
		Id: transaction.ID,
	}, nil
}

func (m *defaultTransactionModel) GetTransactionHistory(ctx context.Context, data *v1alpha1.GetTransactionHistoryRequest) (*v1alpha1.GetTransactionHistoryResponse, error) {
	if data.AccountId == "" {
		log.Println("account is required")
		return nil, errors.New("account is required")
	}

	var result []Transaction
	m.conn.Database().Collection("transaction")
	err := m.conn.Find(ctx, bson.M{
		"$or": []bson.M{
			{"from_account": data.AccountId},
			{"to_account": data.AccountId},
		},
	}, &result)
	if err != nil {
		log.Printf("failed to fetch transaction history for account %s: %v", data.AccountId, err)
		return nil, err
	}

	var transactions []*v1alpha1.Transaction
	for _, t := range result {
		transactions = append(transactions, &v1alpha1.Transaction{
			Id:          t.ID,
			FromAccount: t.FromAccount,
			ToAccount:   t.ToAccount,
			Amount:      float32(t.Amount),
			Timestamp:   t.Timestamp,
		})
	}

	return &v1alpha1.GetTransactionHistoryResponse{
		Transactions: transactions,
	}, nil
}

func (m *defaultTransactionModel) UpdateAccountBalance(ctx context.Context, account *Transaction) (*Transaction, error) {
	account.Timestamp = time.Now().String()

	m.conn.Database().Collection("accounts")
	_, err := m.conn.UpdateOne(ctx, bson.M{"id": account.ID}, bson.M{
		"$set": bson.M{
			"balance":   account.Amount,
			"update_at": account.Timestamp,
		},
	})
	return &Transaction{
		ID:        account.ID,
		Amount:    account.Amount,
		Timestamp: account.Timestamp,
	}, err
}

func (m *defaultTransactionModel) FindOneByName(ctx context.Context, account_id string) (*Transaction, error) {
	var data Transaction
	m.conn.Database().Collection("accounts")
	err := m.conn.FindOne(ctx, bson.M{"id": account_id}, &data)
	if err != nil {
		log.Println(err)
	}
	return &data, nil
}

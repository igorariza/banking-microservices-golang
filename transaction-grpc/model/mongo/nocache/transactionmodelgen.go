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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type transactionModel interface {
	TransferMoney(ctx context.Context, data *v1alpha1.TransferMoneyRequest) (*v1alpha1.TransferMoneyResponse, error)
	GetTransactionHistory(ctx context.Context, data *v1alpha1.GetTransactionHistoryRequest) (*v1alpha1.GetTransactionHistoryResponse, error)
	UpdateAccountBalance(ctx context.Context, account *Account) (*Account, error)
	FindOneByName(ctx context.Context, name string) (*Account, error)
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
	if int64(fromAccount.Balance) < int64(data.Amount) {
		return nil, errors.New("insufficient balance")
	}

	fromAccount.Balance -= float64(data.Amount)
	toAccount.Balance += float64(data.Amount)
	_, err = m.UpdateAccountBalance(ctx, fromAccount)
	if err != nil {
		return nil, err
	}
	_, err = m.UpdateAccountBalance(ctx, toAccount)
	if err != nil {
		return nil, err
	}
	trs := &Transaction{
		ID:          uuid.New().String(),
		FromAccount: data.FromAccount,
		ToAccount:   data.FromAccount,
		Amount:      float64(data.Amount),
		Timestamp:   time.Now().String(),
	}
	_, err = m.CreateTransaction(ctx, trs)

	return &v1alpha1.TransferMoneyResponse{
		Id: trs.ID,
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

func (m *defaultTransactionModel) UpdateAccountBalance(ctx context.Context, account *Account) (*Account, error) {
	account.UpdateAt = time.Now().String()
	filter := bson.D{{"id", account.ID}}
	update := bson.D{{"$set", bson.D{{"balance", account.Balance}}}}
	options := options.Update().SetUpsert(true)
	_, err := m.conn.Database().Collection("accounts").UpdateOne(ctx, filter, update, options)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
	}
	return &Account{
		Balance:  account.Balance,
		UpdateAt: account.UpdateAt,
	}, err
}

func (m *defaultTransactionModel) FindOneByName(ctx context.Context, account_id string) (*Account, error) {
	var data Account
	filter := bson.D{{"id", account_id}}
	err := m.conn.Database().Collection("accounts").FindOne(ctx, filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
	}
	return &data, nil
}

func (m *defaultTransactionModel) CreateTransaction(ctx context.Context, transaction *Transaction) (*Transaction, error) {
	_, err := m.conn.Database().Collection("transactions").InsertOne(ctx, transaction)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return transaction, nil

}

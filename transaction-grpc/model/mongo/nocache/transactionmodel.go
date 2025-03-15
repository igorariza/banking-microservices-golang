package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ TransactionModel = (*customTransactionModel)(nil)

type (
	// PolicyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPolicyModel.
	TransactionModel interface {
		transactionModel
	}

	customTransactionModel struct {
		*defaultTransactionModel
	}
)

// NewPolicyModel returns a model for the mongo.
func NewTransactionModel(url, db, collection string) TransactionModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customTransactionModel{
		defaultTransactionModel: newDefaultTransactionModel(conn),
	}
}

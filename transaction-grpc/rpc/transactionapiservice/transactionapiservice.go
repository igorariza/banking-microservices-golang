// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: transaction_api.proto

package transactionapiservice

import (
	"context"

	"banking-system/transaction-service/rpc/types/transaction/v1alpha1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetTransactionHistoryRequest  = v1alpha1.GetTransactionHistoryRequest
	GetTransactionHistoryResponse = v1alpha1.GetTransactionHistoryResponse
	Transaction                   = v1alpha1.Transaction
	TransferMoneyRequest          = v1alpha1.TransferMoneyRequest
	TransferMoneyResponse         = v1alpha1.TransferMoneyResponse

	TransactionAPIService interface {
		TransferMoney(ctx context.Context, in *TransferMoneyRequest, opts ...grpc.CallOption) (*TransferMoneyResponse, error)
		GetTransactionHistory(ctx context.Context, in *GetTransactionHistoryRequest, opts ...grpc.CallOption) (*GetTransactionHistoryResponse, error)
	}

	defaultTransactionAPIService struct {
		cli zrpc.Client
	}
)

func NewTransactionAPIService(cli zrpc.Client) TransactionAPIService {
	return &defaultTransactionAPIService{
		cli: cli,
	}
}

func (m *defaultTransactionAPIService) TransferMoney(ctx context.Context, in *TransferMoneyRequest, opts ...grpc.CallOption) (*TransferMoneyResponse, error) {
	client := v1alpha1.NewTransactionAPIServiceClient(m.cli.Conn())
	return client.TransferMoney(ctx, in, opts...)
}

func (m *defaultTransactionAPIService) GetTransactionHistory(ctx context.Context, in *GetTransactionHistoryRequest, opts ...grpc.CallOption) (*GetTransactionHistoryResponse, error) {
	client := v1alpha1.NewTransactionAPIServiceClient(m.cli.Conn())
	return client.GetTransactionHistory(ctx, in, opts...)
}

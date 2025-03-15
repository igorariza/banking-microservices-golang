package grpc

import (
	"os"
	"sync"

	"banking-system/transaction-service/rpc/types/transaction/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client v1alpha1.TransactionAPIServiceClient
var doOnce sync.Once
var transactionerviceUri string
var transactionerviceTimeout string

func init() {

	doOnce.Do(func() {
		transactionerviceTimeout = os.Getenv("NOTIFICATION_TRANSACTION_TIMEOUT")
		if transactionerviceTimeout == "" {
			transactionerviceTimeout = "30s"
		}
		transactionerviceUri = os.Getenv("NOTIFICATION_TRANSACTION_URI")
		con, err := grpc.Dial(transactionerviceUri, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		client = v1alpha1.NewTransactionAPIServiceClient(con)
	})
}

//Todo: Add the function to call the gRPC service
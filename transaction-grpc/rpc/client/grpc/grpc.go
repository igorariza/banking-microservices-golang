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
var accounterviceUri string
var accounterviceTimeout string

func init() {

	doOnce.Do(func() {
		accounterviceTimeout = os.Getenv("NOTIFICATION_ACCOUNT_TIMEOUT")
		if accounterviceTimeout == "" {
			accounterviceTimeout = "30s"
		}
		accounterviceUri = os.Getenv("NOTIFICATION_ACCOUNT_URI")
		con, err := grpc.Dial(accounterviceUri, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		client = v1alpha1.NewAccountAPIServiceClient(con)
	})
}

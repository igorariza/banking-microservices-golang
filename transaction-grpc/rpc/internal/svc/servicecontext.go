package svc

import (
	model "banking-system/transaction-service/model/mongo/nocache"
	"banking-system/transaction-service/rpc/internal/config"
	//"k8s.io/client-go/kubernetes"
)

type ServiceContext struct {
	Config config.Config
	DB     model.TransactionModel
}

func NewServiceContext(c config.Config, db model.TransactionModel) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}


package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	model "banking-system/transaction-service/model/mongo/nocache"
	"banking-system/transaction-service/pkg/secrets"
	"banking-system/transaction-service/pkg/utils"
	"banking-system/transaction-service/rpc/internal/config"
	"banking-system/transaction-service/rpc/internal/constants"
	"banking-system/transaction-service/rpc/internal/server"
	"banking-system/transaction-service/rpc/internal/svc"
	"banking-system/transaction-service/rpc/types/transaction/v1alpha1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/transactionapi.yaml", "the config file")

func main() {

	secrets.LoadSecrets()
	//corsMiddleware()
	flag.Parse()
	utils.CreateTopic(os.Getenv("KAFKA_BROKER"), os.Getenv("CREATE_TRANSACTION_TOPIC"), 3, 1)

	mongo_uri := os.Getenv("MONGODB_URI")
	db_name := os.Getenv("MONGODB_DB_NAME")
	collection := constants.COLLECTION

	db := model.NewTransactionModel(mongo_uri, db_name, collection)
	cfg := getServiceConfig()
	ctx := svc.NewServiceContext(*cfg, db)
	startGRPCServer(ctx)

}

func startGRPCServer(ctx *svc.ServiceContext) {
	var g config.ConfigGrpc
	g.Name = "transactionapi.rpc"
	g.ListenOn = os.Getenv("TRANSACTION_API_PORT")
	conf.MustLoad(*configFile, &g)

	s := zrpc.MustNewServer(g.RpcServerConf, func(grpcServer *grpc.Server) {
		v1alpha1.RegisterTransactionAPIServiceServer(grpcServer, server.NewTransactionAPIServiceServer(ctx))

		if ctx.Config.Mode == service.DevMode || ctx.Config.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("%s %s\n", constants.START_SERVER, g.ListenOn)
	s.Start()
}

func corsMiddleware() rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next(w, r)
		}
	}
}

func getServiceConfig() *config.Config {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	return &c
}

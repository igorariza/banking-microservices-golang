package main

import (
	"flag"
	"net/http"
	"os"

	model "banking-system/transaction-service/model/mongo/nocache"
	"banking-system/transaction-service/pkg/secrets"
	"banking-system/transaction-service/rpc/internal/config"
	"banking-system/transaction-service/rpc/internal/constants"
	"banking-system/transaction-service/rpc/internal/handler"
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
	// Load secrets and parse flags
	secrets.LoadSecrets()
	flag.Parse()

	// Retrieve MongoDB connection information from environment variables
	mongo_uri := os.Getenv("MONGODB_URI")
	db_name := os.Getenv("MONGODB_DB_NAME")
	collection := constants.COLLECTION

	// Initialize MongoDB connection and context
	db := model.NewAccountModel(mongo_uri, db_name, collection)
	cfg := getServiceConfig()
	ctx := svc.NewServiceContext(*cfg, db)

	// Start gRPC server in a goroutine
	go func() {
		startGRPCServer(ctx)
	}()

	// Start REST server
	startRESTServer(ctx)
}

// Function to start gRPC server
func startGRPCServer(ctx *svc.ServiceContext) {
	var g config.ConfigGrpc
	g.Name = "transaction.rpc"
	g.ListenOn = ":50051"
	conf.MustLoad(*configFile, &g)

	s := zrpc.MustNewServer(g.RpcServerConf, func(grpcServer *grpc.Server) {
		v1alpha1.RegisterAccountAPIServiceServer(grpcServer, server.NewAccountAPIServiceServer(ctx))

		if ctx.Config.Mode == service.DevMode || ctx.Config.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	log.print(constants.START_SERVER, g.ListenOn)
	s.Start()
}

// Function to start REST server
func startRESTServer(ctx *svc.ServiceContext) {
	c := getServiceConfig()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	
	// Middleware para habilitar CORS
	server.Use(corsMiddleware())

	handler.RegisterHandlers(server, ctx)

	log.print(constants.START_SERVER, c.Host, c.Port)
	server.Start()
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

// Function to load service configuration
func getServiceConfig() *config.Config {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	return &c
}

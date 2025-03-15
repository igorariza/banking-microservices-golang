package main

import (
	"banking-system/account-service/data"
	"banking-system/account-service/routes"
	"banking-system/account-service/utils"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// err := router.SetTrustedProxies([]string{"127.0.0.1", "0.0.0.0", "192.168.1.0/24", os.Getenv("KAFKA_BROKER")})
	// if err != nil {
	// 	log.Fatalf("Error setting trusted proxies %v", err)
	// 	return
	// }

	err := utils.CreateTopic(os.Getenv("KAFKA_BROKER"), os.Getenv("CREATE_ACCOUNT_TOPIC"), 3, 1)
	if err != nil {
		log.Fatalf("Error creating topic kafka%v", err)
		return
	}

	ctx := context.TODO()
	client, err := data.Connect(ctx, os.Getenv("MONGODB_URI"))
	database := client.Database(os.Getenv("MONGODB_DB_NAME"))
	defer client.Disconnect(ctx)

	if err != nil {
		log.Fatalf("Error cannot connect to mongodb %v", err)
		return
	}
	routes.SetupAccountRoutes(router, database)

	if err := router.Run(os.Getenv("PORT")); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

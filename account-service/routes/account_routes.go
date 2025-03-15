package routes

import (
	"banking-system/account-service/controllers"
	"banking-system/account-service/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAccountRoutes(router *gin.Engine, db *mongo.Database) {
	accountService := services.NewAccountService(db)
	accountController := controllers.NewAccountController(accountService)
	router.POST("/accounts", accountController.CreateAccount)
	router.GET("/accounts/:id", accountController.GetAccountBalance)
}

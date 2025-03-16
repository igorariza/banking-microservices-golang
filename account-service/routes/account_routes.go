package routes

import (
	"banking-system/account-service/controllers"
	"banking-system/account-service/services"
	"banking-system/account-service/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAccountRoutes(router *gin.Engine, db *mongo.Database) {
	accountService := services.NewAccountService(db)
	accountController := controllers.NewAccountController(accountService)

	router.POST("/generate_token", accountController.GenerateToken)
	protected := router.Group("/accounts")
	protected.Use(utils.JWTMiddleware())
	{
		protected.POST("", accountController.CreateAccount)
		protected.GET("/:id", accountController.GetAccountBalance)
	}
}

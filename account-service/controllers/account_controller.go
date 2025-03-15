package controllers

import (
	"banking-system/account-service/data/models"
	"banking-system/account-service/services"
	"banking-system/account-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService *services.AccountService
}

func NewAccountController(accountService *services.AccountService) *AccountController {
	return &AccountController{accountService: accountService}
}

func (c *AccountController) CreateAccount(ctx *gin.Context) {
	var account models.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	validate_account := utils.ValidateCreateAccount(account)
	if validate_account != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validate_account.Error()})
		return
	}

	utils.GenerateToken(account.Name)

	createdAccount, err := c.accountService.CreateAccount(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdAccount)
}

func (c *AccountController) GetAccountBalance(ctx *gin.Context) {
	id := ctx.Param("id")
	balance, err := c.accountService.GetAccountBalance(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"balance": balance})
}

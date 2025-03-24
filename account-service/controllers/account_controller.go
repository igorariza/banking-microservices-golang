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
	var account *models.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	validateChan := make(chan error)
	go func() {
		validateChan <- utils.ValidateCreateAccount(*account)
	}()

	createChan := make(chan *models.Account)
	errorChan := make(chan error)
	go func() {
		createdAccount, err := c.accountService.CreateAccount(*account)
		if err != nil {
			errorChan <- err
			return
		}
		createChan <- createdAccount
	}()
	validateErr := <-validateChan
	if validateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
		return
	}
	var createdAccount *models.Account
	select {
	case createdAccount = <-createChan:
	case err := <-errorChan:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdAccount)
}

func (c *AccountController) GetAccountBalance(ctx *gin.Context) {
	id := ctx.Param("id")
	balanceChan := make(chan float64)
	errorChan := make(chan error)
	go func() {
		select {
		case <-ctx.Done():
			errorChan <- ctx.Err()
		default:
			balance, err := c.accountService.GetAccountBalance(id)
			if err != nil {
				errorChan <- err
				return
			}
			balanceChan <- balance
		}
	}()

	select {
	case balance := <-balanceChan:
		ctx.JSON(http.StatusOK, gin.H{"balance": balance})
	case err := <-errorChan:
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found", "details": err.Error()})
	}
}

func (c *AccountController) GenerateToken(ctx *gin.Context) {
	var account models.Account
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user name"})
		return
	}
	tokenChan := make(chan string)
	errorChan := make(chan error)
	go func() {
		token, err := utils.GenerateToken(account.Name)
		if err != nil {
			errorChan <- err
			return
		}
		tokenChan <- token
	}()
	select {
	case token := <-tokenChan:
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	case err := <-errorChan:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token", "details": err.Error()})
	}
}

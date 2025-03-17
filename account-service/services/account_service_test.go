package services_test

import (
	"banking-system/account-service/data/models"
	"banking-system/account-service/services"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateAccount_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Create Account Successfully", func(mt *mtest.T) {
		service := services.NewAccountService(mt.DB)
		account := models.Account{
			Name:    "Test Account",
			Balance: 100.0,
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		createdAccount, err := service.CreateAccount(account)
		fmt.Print(err)
		fmt.Print(createdAccount)
	})
}
func TestCreateAccount_AlreadyExists(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Account Already Exists", func(mt *mtest.T) {
		service := services.NewAccountService(mt.DB)
		account := models.Account{
			ID:      uuid.New().String(),
			Name:    "Existing Account",
			Balance: 200.0,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "accounts", mtest.FirstBatch, bson.D{{"name", "Existing Account"}}))
		_, err := service.CreateAccount(account)
		assert.Error(t, err)
	})
}

func TestGetAccountBalance_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Get Account Balance Successfully", func(mt *mtest.T) {
		service := services.NewAccountService(mt.DB)
		account := models.Account{
			ID:      uuid.New().String(),
			Name:    "Test Account",
			Balance: 500.0,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "accounts", mtest.FirstBatch, bson.D{{"id", account.ID}, {"balance", account.Balance}}))
		balance, _ := service.GetAccountBalance(account.ID)
		assert.Equal(t, 0, int(balance))
	})
}

func TestGetAccountBalance_NotFound(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Account Not Found", func(mt *mtest.T) {
		service := services.NewAccountService(mt.DB)

		mt.AddMockResponses(mtest.CreateCursorResponse(0, "accounts", mtest.FirstBatch))
		_, err := service.GetAccountBalance("non-existing-id")
		assert.Error(t, err)
	})
}

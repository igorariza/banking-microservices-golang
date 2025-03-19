package services

import (
	"banking-system/account-service/data"
	"banking-system/account-service/data/models"
	"banking-system/account-service/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type AccountService struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewAccountService(db *mongo.Database) *AccountService {
	return &AccountService{
		Client:   db.Client(),
		Database: db,
	}
}

func (s *AccountService) CreateAccount(account models.Account) (*models.Account, error) {
	create, err := data.CreateAccount(s.Client)
	if err != nil {
		return nil, err
	}
	utils.PublishAccountEvent(context.Background(), account.ID)

	return &models.Account{
		ID:       create.ID,
		Name:     create.Name,
		Balance:  create.Balance,
		CreateAt: create.CreateAt,
		UpdateAt: create.UpdateAt,
	}, nil
}

func (s *AccountService) GetAccountBalance(id string) (float64, error) {
	var account models.Account
	account.ID = id
	_, err := data.GetAccountBalance(s.Client, account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, errors.New("account not found")
		}
		return 0, err
	}

	utils.PublishAccountEvent(context.Background(), account.ID)

	return account.Balance, nil
}

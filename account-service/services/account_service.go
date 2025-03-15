package services

import (
	"banking-system/account-service/data/models"
	"banking-system/account-service/utils"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
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

func validateIfExistAccountName(s *AccountService, name string) (bool, error) {
	var account models.Account
	account_name := s.Database.Collection("accounts")
	filter := bson.M{"name": name}
	err := account_name.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *AccountService) CreateAccount(account models.Account) (*models.Account, error) {

	if exist, err := validateIfExistAccountName(s, account.Name); exist {
		return nil, errors.New("account name already exists")
	} else if err != nil {
		return nil, err
	}

	account.ID = uuid.New().String()
	account.CreateAt = time.Now().String()
	account.UpdateAt = time.Now().String()

	_, err := s.Client.Database(s.Database.Name()).Collection("accounts").InsertOne(context.Background(), account)
	if err != nil {
		return nil, err
	}
	utils.PublishAccountEvent(context.Background(), account.ID)
	

	return &models.Account{
		ID:       account.ID,
		Name:     account.Name,
		Balance:  account.Balance,
		CreateAt: account.CreateAt,
		UpdateAt: account.UpdateAt,
	}, nil
}

func (s *AccountService) GetAccountBalance(id string) (float64, error) {
	var account models.Account

	err := s.Client.Database(s.Database.Name()).Collection("accounts").FindOne(context.Background(), bson.M{"id": id}).
		Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, errors.New("account not found")
		}
		return 0, err
	}

	utils.PublishAccountEvent(context.Background(), account.ID)
	
	return account.Balance, nil
}

func (s *AccountService) TransferMoney(fromAccountID string, toAccountID string, amount float64) (string, error) {
	if amount < 0 {
		return "", errors.New("amount cannot be negative")
	}

	fromAccount, err := s.GetAccountBalance(fromAccountID)
	if err != nil {
		return "", err
	}

	toAccount, err := s.GetAccountBalance(toAccountID)
	if err != nil {
		return "", err
	}

	if fromAccount < amount {
		return "", errors.New("insufficient balance")
	}

	fromAccount -= amount
	toAccount += amount

	_, err = s.Client.Database(s.Database.Name()).Collection("accounts").UpdateOne(context.Background(), bson.M{"id": fromAccountID},
		bson.M{"$set": bson.M{"balance": fromAccount}})
	if err != nil {
		return "", err
	}

	_, err = s.Client.Database(s.Database.Name()).Collection("accounts").UpdateOne(context.Background(), bson.M{"id": toAccountID},
		bson.M{"$set": bson.M{"balance": toAccount}})
	if err != nil {
		return "", err
	}

	return uuid.New().String(), nil
}

package data

import (
	"banking-system/account-service/data/models"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func validateIfExistAccountName(db *mongo.Database, name string) (bool, error) {
	var account models.Account
	account_name := db.Collection("accounts")
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

func CreateAccount(client *mongo.Client) (*models.Account, error) {
	var account models.Account
	db := client.Database("banking-service")

	if exist, err := validateIfExistAccountName(db, account.Name); exist {
		return nil, errors.New("account name already exists")
	} else if err != nil {
		return nil, err
	}

	account.ID = uuid.New().String()
	account.CreateAt = time.Now().String()
	account.UpdateAt = time.Now().String()
	_, err := db.Collection("accounts").InsertOne(context.Background(), account)
	if err != nil {
		return nil, err
	}

	return &models.Account{
		ID:       account.ID,
		Name:     account.Name,
		Balance:  account.Balance,
		CreateAt: account.CreateAt,
		UpdateAt: account.UpdateAt,
	}, nil
}
func GetAccountBalance(client *mongo.Client, ac models.Account) (float64, error) {
	var account models.Account
	account.ID = ac.ID
	filter := bson.D{{"id", account.ID}}
	db := client.Database("banking-service")
	err := db.Collection("accounts").FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, errors.New("account not found")
		}
		return 0, err
	}

	return account.Balance, nil
}

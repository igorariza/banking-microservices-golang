package data

import (
	"banking-system/account-service/data/models"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func validateIfExistAccountName(db *mongo.Database, name string) (bool, error) {
	var account models.Account
	account_name := db.Collection("accounts")
	filter := bson.D{{"name", name}}
	err := account_name.FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateAccount(client *mongo.Client, ac *models.Account) (*models.Account, error) {
	var account models.Account
	account.Name = ac.Name
	db := client.Database(os.Getenv("MONGODB_DB_NAME"))

	if exist, err := validateIfExistAccountName(db, account.Name); exist {
		return nil, errors.New("account name already exists")
	} else if err != nil {
		return nil, err
	}

	account.ID = uuid.New().String()
	account.Name = ac.Name
	account.CreateAt = time.Now().String()
	account.UpdateAt = time.Now().String()
	account.Balance = ac.Balance
	_, err := db.Collection("accounts").InsertOne(context.Background(), account)
	if err != nil {
		return nil, err
	}

	return &models.Account{
		ID:       account.ID,
		Name:     ac.Name,
		Balance:  ac.Balance,
		CreateAt: account.CreateAt,
		UpdateAt: account.UpdateAt,
	}, nil
}
func GetAccountBalance(client *mongo.Client, ac models.Account) (float64, error) {
	var account models.Account
	fmt.Println("ac.Id", ac.ID)
	filter := bson.D{{"id", ac.ID}}
	db := client.Database(os.Getenv("MONGODB_DB_NAME"))
	err := db.Collection("accounts").FindOne(context.Background(), filter).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}
		return 0, err
	}
	fmt.Println("account.Balance", account.Balance)

	return account.Balance, nil
}

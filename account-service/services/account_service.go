package services

import (
	"banking-system/account-service/data"
	"banking-system/account-service/data/models"
	"banking-system/account-service/utils"
	"context"
	"errors"
	"log"
	"time"

	"github.com/sony/gobreaker"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountService struct {
	Client   *mongo.Client
	Database *mongo.Database
	Cb       *gobreaker.CircuitBreaker
}

func NewAccountService(db *mongo.Database) *AccountService {

	cbSettings := gobreaker.Settings{
		Name:        "AccountServiceCircuitBreaker",
		MaxRequests: 5,
		Interval:    time.Second * 60,
		Timeout:     time.Second * 10,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 5 && failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("State Change: %s: %s -> %s", name, from, to)
		},
	}
	cb := gobreaker.NewCircuitBreaker(cbSettings)

	return &AccountService{
		Client:   db.Client(),
		Database: db,
		Cb:       cb,
	}
}

func (s *AccountService) CreateAccount(account models.Account) (*models.Account, error) {
	acc, err := s.Cb.Execute(func() (interface{}, error) {
		create, err := data.CreateAccount(s.Client, &account)
		if err != nil {
			return nil, err
		}
		utils.PublishAccountEvent(context.Background(), account.ID)

		return create, nil
	})

	if err != nil {
		log.Printf("Error creating account: %v", err)
		return nil, errors.New(err.Error())
	}

	return acc.(*models.Account), nil
}

func (s *AccountService) GetAccountBalance(id string) (float64, error) {
	var account models.Account
	account.ID = id
	result, err := s.Cb.Execute(func() (interface{}, error) {
		bl, err := data.GetAccountBalance(s.Client, account)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return 0, errors.New(err.Error())
			}
			return 0, err
		}
		account.Balance = bl

		utils.PublishAccountEvent(context.Background(), account.ID)
		return account, nil
	})

	if err == nil {
		account = result.(models.Account)
	}

	if err != nil {
		log.Printf("Error getting account balance: %v", err)
		return 0, errors.New("service unavailable")
	}

	return account.Balance, nil
}

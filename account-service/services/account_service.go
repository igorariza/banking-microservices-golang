package services

import (
	"banking-system/account-service/data"
	"banking-system/account-service/data/models"
	"banking-system/account-service/utils"
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/sirupsen/logrus"
	"github.com/sony/gobreaker"
	"go.mongodb.org/mongo-driver/mongo"
	elogrus "gopkg.in/sohlich/elogrus.v7"
)

type AccountService struct {
	Client   *mongo.Client
	Database *mongo.Database
	Cb       *gobreaker.CircuitBreaker
	DataDog  *statsd.Client
	Logger   *logrus.Logger
}

func NewAccountService(db *mongo.Database) *AccountService {

	statsdAddress := os.Getenv("STATSD_URL")
	if statsdAddress == "" {
		statsdAddress = "127.0.0.1:8125"
	}
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

	statsdClient, err := statsd.New(statsdAddress)
	if err != nil {
		log.Fatalf("Error creating datadog client: %v", err)
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	elasticClient := utils.NewElasticClient()
	hook, err := elogrus.NewAsyncElasticHook(elasticClient, "localhost", logrus.DebugLevel, "account-service")
	if err != nil {
		log.Fatalf("Error creating elastic hook: %v", err)
	}
	logger.AddHook(hook)

	return &AccountService{
		Client:   db.Client(),
		Database: db,
		Cb:       cb,
		DataDog:  statsdClient,
	}
}

func (s *AccountService) CreateAccount(account models.Account) (*models.Account, error) {
	start := time.Now()

	acc, err := s.Cb.Execute(func() (interface{}, error) {
		create, err := data.CreateAccount(s.Client, &account)
		if err != nil {
			return nil, err
		}
		utils.PublishAccountEvent(context.Background(), account.ID)

		return create, nil
	})

	duration := time.Since(start)
	s.DataDog.Timing("account.create", duration, []string{"success:true"}, 1)
	if err != nil {
		s.DataDog.Incr("account.create", []string{"success:false"}, 1)
		s.Logger.WithFields(logrus.Fields{
			"operation": "create_account",
			"success":   false,
			"duration":  duration,
			"error":     err.Error(),
		}).Error("Failed to create account")
		return nil, errors.New("service unavailable")
	}

	s.DataDog.Incr("account.create", []string{"success:true"}, 1)
	s.Logger.WithFields(logrus.Fields{
		"operation":  "create_account",
		"success":    true,
		"duration":   duration,
		"account_id": acc.(*models.Account).ID,
	}).Info("Account created successfully")
	return acc.(*models.Account), nil
}

func (s *AccountService) GetAccountBalance(id string) (float64, error) {
	start := time.Now()
	var account models.Account
	account.ID = id
	result, err := s.Cb.Execute(func() (interface{}, error) {
		bl, err := data.GetAccountBalance(s.Client, account)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return 0, errors.New("account not found")
			}
			return 0, err
		}
		account.Balance = bl

		utils.PublishAccountEvent(context.Background(), account.ID)
		return account, nil
	})

	duration := time.Since(start)
	s.DataDog.Timing("account.get_balance", duration, []string{"success:true"}, 1)

	if err == nil {
		account = result.(models.Account)
	}

	if err != nil {
		s.DataDog.Incr("account.get_balance", []string{"success:false"}, 1)
		s.Logger.WithFields(logrus.Fields{
			"operation": "get_balance",
			"success":   false,
			"duration":  duration,
			"error":     err.Error(),
		}).Error("Failed to get account balance")
		return 0, errors.New("service unavailable")
	}

	s.DataDog.Incr("account.get_balance", []string{"success:true"}, 1)
	s.Logger.WithFields(logrus.Fields{
		"operation":  "get_balance",
		"success":    true,
		"duration":   duration,
		"account_id": account.ID,
		"balance":    account.Balance,
	}).Info("Account balance retrieved successfully")

	return account.Balance, nil
}

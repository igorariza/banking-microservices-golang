package model

import (
	"time"
)

type Transaction struct {
	ID          int       `json:"id" bson:"_id,omitempty"`
	FromAccount int64       `json:"from_account" bson:"from_account"`
	ToAccount   int64       `json:"to_account" bson:"to_account"`
	Amount      float64   `json:"amount" bson:"amount"`
	Timestamp   time.Time `json:"timestamp" bson:"timestamp"`
}

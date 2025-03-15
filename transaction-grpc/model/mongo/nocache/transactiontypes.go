package model

import (
	"time"
)

type Transaction struct {
	ID          int       `json:"id" bson:"_id,omitempty"`
	FromAccount string       `json:"from_account" bson:"from_account"`
	ToAccount   string       `json:"to_account" bson:"to_account"`
	Amount      float64   `json:"amount" bson:"amount"`
	Timestamp   time.Time `json:"timestamp" bson:"timestamp"`
}

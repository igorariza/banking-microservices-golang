package model

type Transaction struct {
	ID          int     `json:"id" bson:"_id,omitempty"`
	FromAccount string  `json:"from_account" bson:"from_account"`
	ToAccount   string  `json:"to_account" bson:"to_account"`
	Amount      float64 `json:"amount" bson:"amount"`
	Timestamp   string  `json:"timestamp" bson:"timestamp"`
}

package model

type Transaction struct {
	ID          string     `json:"id" bson:"id,omitempty"`
	FromAccount string  `json:"from_account" bson:"from_account"`
	ToAccount   string  `json:"to_account" bson:"to_account"`
	Amount      float64 `json:"amount" bson:"amount"`
	Timestamp   string  `json:"timestamp" bson:"timestamp"`
}

type Account struct {
	ID       string  `bson:"id,omitempty" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Balance  float64 `bson:"balance" json:"balance"`
	CreateAt string  `bson:"create_at" json:"create_at"`
	UpdateAt string  `bson:"update_at" json:"update_at"`
}

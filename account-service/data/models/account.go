package models

type Account struct {
	ID       string  `bson:"id,omitempty" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Balance  float64 `bson:"balance" json:"balance"`
	CreateAt string  `bson:"create_at" json:"create_at"`
	UpdateAt string  `bson:"update_at" json:"update_at"`
}

type Transaction struct {
	ID          string  `bson:"_id,omitempty" json:"id"`
	FromAccount string  `bson:"from_account" json:"from_account"`
	ToAccount   string  `bson:"to_account" json:"to_account"`
	Amount      float64 `bson:"amount" json:"amount"`
	CreateAt    string  `bson:"create_at" json:"create_at"`
	UpdateAt    string  `bson:"update_at" json:"update_at"`
}

type TransactionResponse struct {
	TransactionID string `json:"transaction_id"`
}

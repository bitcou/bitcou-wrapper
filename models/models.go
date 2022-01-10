package models

import "github.com/shopspring/decimal"

type CreateOrderEncryptedInput struct {
	TransactionId string
	TxId          string
	TotalValue    decimal.Decimal
	UserInfo      []byte
}

type FirebaseAccount struct {
	Address   string
	Purchases []FirebasePurchase
}

type FirebasePurchase struct {
	ID            string
	ProductId     string
	TotalValue    float64
	TransactionId string
	Status        string
	Timestamp     int64
}

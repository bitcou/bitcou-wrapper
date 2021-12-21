package models

import "github.com/shopspring/decimal"

type CreateOrderEncryptedInput struct {
	TransactionId string
	TxId          string
	TotalValue    decimal.Decimal
	UserInfo      string
}

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
	ID                 string  `json:"id"`
	ProductId          string  `json:"product_id"`
	TotalValue         float64 `json:"total_value"`
	TransactionId      string  `json:"transaction_id"`
	Status             string  `json:"status"`
	Timestamp          int64   `json:"timestamp"`
	Code               string  `json:"code"`
	RedeemInstructions string  `json:"redeem_instructions"`
	RedeemUrl          string  `json:"redeem_url"`
}
type FirebaseNonce struct {
	Nonce int `json:"nonce"`
}

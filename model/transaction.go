package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID            uint
	ContractNumber    string
	OTR               float64
	AdminFee          float64
	NumOfInstallments int
	LeftInstallments  int
	InterestRate      float64
	AssetName         string
	ProductLink       string
}

type TransactionModel interface {
	CreateTransaction(transaction Transaction) (Transaction, error)
	GetTransactionsByUserID(userID uint) ([]Transaction, error)
}

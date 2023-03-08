package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CustomerID        uint
	ContractNumber    string
	OTR               float64
	AdminFee          float64
	NumOfInstallments int
	LeftInstallments  int
	InterestRate      float64
	AssetName         string
}

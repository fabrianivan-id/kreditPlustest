package models

import "gorm.io/gorm"

type LoanRequest struct {
	gorm.Model
	UserID          uint
	RequestedAmount int
	Status          uint
}

type LoanRequestModel interface {
	CreateLoanRequest(request LoanRequest) (LoanRequest, error)
	GetLoanRequestsByUserID(userID uint) ([]LoanRequest, error)
	GetLoanRequest(requestID uint) (LoanRequest, error)
	UpdateLoanRequest(request LoanRequest) (LoanRequest, error)
}

package models

import (
	"gorm.io/gorm"
	"time"
)

type LoanLimit struct {
	gorm.Model
	UserID   uint
	Limit    int
	Duration time.Duration
	Status   uint
}

type LoanLimitModel interface {
	CreateLoanLimit(limit LoanLimit) (LoanLimit, error)
	GetLoanLimit(userID uint) (LoanLimit, error)
	UpdateLoanLimit(limit LoanLimit) (LoanLimit, error)
}

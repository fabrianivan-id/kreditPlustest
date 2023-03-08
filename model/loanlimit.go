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
}

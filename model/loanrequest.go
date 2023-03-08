package models

import "gorm.io/gorm"

type LoanRequest struct {
	gorm.Model
	UserID          uint `json:"token" form:"token"`
	RequestedAmount int
	Status          string
}

package models

import (
	"kreditPlus-test/middlewares"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	PhoneNumber  string        `json:"PhoneNumber" form:"PhoneNumber"`
	Password     string        `json:"password" form:"password"`
	NIK          string        `json:"NIK" form:"NIK"`
	FullName     string        `json:"FullName" form:"FullName"`
	LegalName    string        `json:"LegalName" form:"LegalName"`
	BirthPlace   string        `json:"BirthPlace" form:"BirthPlace"`
	BirthDate    string        `json:"BirthDate" form:"BirthDate"`
	Salary       int           `json:"Salary" form:"Salary"`
	KTPImage     []byte        `json:"KTPImage" form:"KTPImage"`
	Selfie       []byte        `json:"Selfie" form:"Selfie"`
	Token        string        `json:"token" form:"token"`
	LoanRequests []LoanRequest `gorm:"foreignKey:UserID"`
	LoanLimits   []LoanLimit   `gorm:"foreignKey:UserID"`
	Transactions []Transaction
}

type GormCustomerModel struct {
	db *gorm.DB
}

func NewCustomerModel(db *gorm.DB) *GormCustomerModel {
	return &GormCustomerModel{db: db}
}

type CustomerModel interface {
	Register(Customer) (Customer, error)
	Login(no_hp, password string) (Customer, error)
	GetUserData(userId int) (Customer, error)
}

func (m *GormCustomerModel) Register(customer Customer) (Customer, error) {
	// Encrypt Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	if err != nil {
		return customer, err
	}
	customer.Password = string(hashedPassword)

	if err := m.db.Save(&customer).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

func (m *GormCustomerModel) Login(email, password string) (Customer, error) {
	var customer Customer
	var err error

	if err = m.db.Where("email = ?", email).First(&customer).Error; err != nil {
		return customer, err
	}

	// Start Checking Hash Password
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if err != nil {
		return customer, err
	}

	customer.Token, err = middlewares.CreateToken(int(customer.ID))

	if err != nil {
		return customer, err
	}

	if err := m.db.Save(customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

func (m *GormCustomerModel) GetUserData(userId int) (Customer, error) {
	var customer Customer
	if err := m.db.First(&customer, userId).Error; err != nil {
		return customer, err
	}
	return customer, nil
}

package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"kreditPlus-test/config"
)

const TimeLayout = "2006-01-02 15:04:05"

var (
	DBLOAN *gorm.DB
)

func Init() {
	var err error

	// DB QR MPM
	url := config.LoanDatabaseUrl
	usr := config.LoanDatabaseUser
	pw := config.LoanDatabasePassword
	name := config.LoanDatabaseName

	DBLOAN, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True", usr, pw, url, name))
	if err != nil {
		panic(err)
	}

	DBLOAN.SingularTable(true)
	DBLOAN.LogMode(true)

}

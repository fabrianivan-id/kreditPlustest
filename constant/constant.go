package constant

import "time"

var (
	// Server
	ServerName         = "kreditPlus"
	ServerPort         = ":2023"
	ServerDefaultRoute = "/api/v1"
	SECRET_JWT         = "RAHASIA"

	// Vikendi
	VikendiBaseURL = ""
	VikendiTimeout = 20 * time.Second

	// KAFKA
	KafkaURL          = ""
	KafkaTopicSuccess = ""

	// REDIS
	//RedisURL = "redis-brimo:6379"
	//RedisURL = "http://172.18.136.92:6379/"

	// Other
	NilStruct          struct{}
	CtxRequestStr      = "request"
	TblProductsLoan    = "tbl_products_loan"
	Custom             = "custom"
	ErrorHandling      = "EH"
	ErrorHandlingDesc  = "Error Handlin"
	TblListPayment     = "tbl_user"
	TblTrxPaymentBrimo = "tbl_trx_user"
	IdParameter        = "KREDITPLUS_LOAN"
	TblParameter       = "tbl_parameters"
)

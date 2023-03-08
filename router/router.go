package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"kreditPlus-test/constant"
)

func Route(
	e *echo.Echo,
	customerController *controller.CustomerController,
	loanReqController *controller.LoanController,
) {
	// ------------------------------------------------------------------
	// Auth Login & Register
	// ------------------------------------------------------------------
	e.POST("/api/register", customerController.RegisterCustomerController)
	e.POST("/api/login", customerController.LogincustomerController)

	// ------------------------------------------------------------------
	// Auth JWT
	// ------------------------------------------------------------------
	jwtMiddleware := middleware.JWT([]byte(constant.SECRET_JWT))

	// ------------------------------------------------------------------
	// Layanan Loan Request
	// ------------------------------------------------------------------
	e.GET("/api/loanrequest", controller.GetAllRecipeController, jwtMiddleware)

	// ------------------------------------------------------------------
	// Loan Limit
	// ------------------------------------------------------------------
	e.GET("/api/categories", categoryController.GetAllCategoryController, jwtMiddleware)
	e.POST("/api/categories", customerControllerontroller.InsertCategoryController, jwtMiddleware)
	e.GET("/api/categories/:id", categoryController.GetCategoryController, jwtMiddleware)
	e.PUT("/api/categories/:id", categoryController.EditCategoryController, jwtMiddleware)
	e.DELETE("/api/categories/:id", categoryController.DeleteCategoryController, jwtMiddleware)

}

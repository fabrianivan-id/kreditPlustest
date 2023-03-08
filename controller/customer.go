package auth

import (
	echo "github.com/labstack/echo/v4"
	"kreditPlus-test/model"
	"net/http"
)

// ------------------------------------------------------------------
// Start Request
// ------------------------------------------------------------------

type CustomerRequest struct {
	NIK        string `json:"NIK" form:"NIK"`
	FullName   string `json:"FullName" form:"FullName"`
	LegalName  string `json:"LegalName" form:"LegalName"`
	BirthPlace string `json:"BirthPlace" form:"BirthPlace"`
	BirthDate  string `json:"BirthDate" form:"BirthDate"`
	Salary     int    `json:"Salary" form:"Salary"`
	KTPImage   []byte `json:"KTPImage" form:"KTPImage"`
	Selfie     []byte `json:"Selfie" form:"Selfie"`
}

type LoginCustomerRequest struct {
	PhoneNumber string `json:"PhoneNumber" form:"PhoneNumber"`
	Password    string `json:"password" form:"password"`
}
type RegisterCustomerRequest struct {
	PhoneNumber string `json:"PhoneNumber" form:"PhoneNumber"`
	Password    string `json:"password" form:"password"`
}


// ------------------------------------------------------------------
// End Request
// ------------------------------------------------------------------

type CustomerController struct {
	customerModel model.CustomerModel
}

func NewCustomerController(userModel model.CustomerModel) *CustomerController {
	return &CustomerController{
		customerModel,
	}
}

func (controller *CustomerController) RegisterUserController(c echo.Context) error {
	var customerRequest RegisterCustomerRequest

	c.Bind(&customerRequest)
	/*
		getLocation, errLocation := location.Geocoding(userRequest.Address)
		if errLocation != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"success": false,
				"code":    500,
				"message": "Internal Server Error",
			})
		}
	*/
	customer := models.Customer{
		PhoneNumber: customerRequest.
	}
	/*if user.Name == "" || user.Email == "" || user.Password == "" || user.Address == "" || user.Gender == "" || user.Role == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}
	*/
	_, err := controller.customerModel.Register(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"code":    500,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Register Account",
	})
}

func (controller *CustomerController) LoginUserController(c echo.Context) error {
	var userRequest LoginCustomerRequest

	c.Bind(&userRequest)

	data, err := controller.customerModel.Login(userRequest.PhoneNumber, userRequest.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Login",
		"token":   data.Token,
	})
}

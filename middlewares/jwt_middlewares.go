package middlewares

import (
	"kreditPlus-test/constant"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = int(userId)
	//claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}

func LoggerMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))
}

func ExtractTokenUser(c echo.Context) (uint, string) {

	// fmt.Println(c.Get("user"))
	// if temp := c.Get("user"); temp != nil {

	// 	u := temp.(*jwt.Token)
	// 	claims := u.Claims.(jwt.MapClaims)
	// 	userId := int(claims["userId"].(float64))
	// 	role := claims["role"].(string)
	// 	return userId, role
	// }
	// return 0, ""

	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userId := uint(claims["userId"].(float64))
		role := claims["role"].(string)
		return userId, role
	}
	return 0, ""
}

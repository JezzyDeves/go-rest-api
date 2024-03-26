package routes

import (
	"net/http"
	"time"

	"github.com/JezzyDeves/go-rest-api/api"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var SecretKey = []byte("secret key")

func InitLoginRoute(e *echo.Echo) {
	api.Echo.GET("/login", func(c echo.Context) error {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		})

		tokenString, _ := token.SignedString(SecretKey)

		return c.JSON(http.StatusOK, tokenString)
	})
}

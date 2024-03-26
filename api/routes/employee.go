package routes

import (
	"net/http"

	"github.com/JezzyDeves/go-rest-api/db"
	"github.com/JezzyDeves/go-rest-api/db/models"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitEmployeeRoutes(e *echo.Echo) {
	g := e.Group("/employees")

	g.Use(echojwt.JWT(SecretKey))

	g.GET("", func(c echo.Context) error {
		var employees []models.Employee

		db.Database.Find(&employees)

		return c.JSON(http.StatusOK, employees)
	})

}

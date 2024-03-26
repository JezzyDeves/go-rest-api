package routes

import (
	"net/http"
	"time"

	"github.com/JezzyDeves/go-rest-api/db"
	"github.com/JezzyDeves/go-rest-api/db/models"
	"github.com/JezzyDeves/go-rest-api/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeCreateRequest struct {
	Name        string    `json:"name" validate:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	JobTitle    string    `json:"jobTitle" validate:"required"`
	Salary      float32   `json:"salary" validate:"required"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
}

func InitEmployeeRoutes(e *echo.Echo) {
	g := e.Group("/employees")

	g.Use(echojwt.JWT(SecretKey))

	g.GET("", func(c echo.Context) error {
		var employees []models.Employee

		db.Database.Find(&employees)

		return c.JSON(http.StatusOK, employees)
	})

	g.POST("/create", func(c echo.Context) error {
		var employee EmployeeCreateRequest

		err := c.Bind(&employee)

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		err = validator.Validate.Struct(&employee)

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)

		db.Database.Create(&models.Employee{
			Name:         employee.Name,
			DateOfBirth:  employee.DateOfBirth,
			JobTitle:     employee.JobTitle,
			Salary:       employee.Salary,
			Username:     employee.Username,
			PasswordHash: string(passwordHash),
		})

		return c.NoContent(http.StatusCreated)
	})
}

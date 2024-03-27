package routes

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/JezzyDeves/go-rest-api/db"
	"github.com/JezzyDeves/go-rest-api/db/models"
	"github.com/JezzyDeves/go-rest-api/validator"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

		var existingUser []models.Employee
		db.Database.Where(&models.Employee{Username: employee.Username}).Find(&existingUser)

		if len(existingUser) > 0 {
			return c.String(http.StatusBadRequest, fmt.Sprintf("A user with the username %v already exists", employee.Username))
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

	type EmployeeLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type EmployeeLoginClaims struct {
		jwt.RegisteredClaims
		Username string
	}

	g.POST("/login", func(c echo.Context) error {
		var loginInfo EmployeeLoginRequest
		c.Bind(&loginInfo)

		var employee models.Employee
		err := db.Database.
			Where(&models.Employee{Username: loginInfo.Username}).
			First(&employee).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.String(http.StatusNotFound, "No user found")
		}

		err = bcrypt.CompareHashAndPassword([]byte(employee.PasswordHash), []byte(loginInfo.Password))

		if err != nil {
			return c.String(http.StatusBadRequest, "Incorrect login information")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, EmployeeLoginClaims{
			Username: loginInfo.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			},
		})

		tokenString, _ := token.SignedString(SecretKey)

		return c.String(http.StatusOK, tokenString)
	})
}

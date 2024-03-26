package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string    `json:"name"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	JobTitle     string    `json:"jobTitle"`
	Salary       float32   `json:"salary"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"passwordHash"`
}

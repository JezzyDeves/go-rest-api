package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string
	DateOfBirth  time.Time
	JobTitle     string
	Salary       float32
	Username     string
	PasswordHash string
}

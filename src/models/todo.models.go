package models

import (
	"gorm.io/gorm"
)

// Defines todo table for database communication
type Todo struct {
	gorm.Model
	Name        string
	Description string
}

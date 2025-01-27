package entities

import (
	"time"

	"gorm.io/gorm"
)

type Holiday struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedBy   uint      `json:"created_by"`
}

package entities

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EventType   string    `json:"type"`
	Start       string    `json:"start"`
	End         string    `json:"end"`
	Date        time.Time `json:"date"`
	CreatedBy   uint      `json:"created_by"`
}

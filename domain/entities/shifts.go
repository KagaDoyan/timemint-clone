package entities

import "gorm.io/gorm"

type Shift struct {
	gorm.Model
	Name        string `gorm:"size:100;not null"`
	Description string `gorm:"type:text"`
	StartTime   string `gorm:"not null"` // Shift start time
	EndTime     string `gorm:"not null"` // Shift end time
	Color       string
}

package entities

import "gorm.io/gorm"

type Shift struct {
	gorm.Model
	Name        string `gorm:"size:100;not null"`
	ShiftType   string `gorm:"type:enum('Permanent', 'One-Time');not null"` // Shift type
	StartTime   string `gorm:"not null"`                                    // Shift start time
	EndTime     string `gorm:"not null"`                                    // Shift end time
	Description string `gorm:"type:text"`
}

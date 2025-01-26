package entities

import "gorm.io/gorm"

//for allow to check in and check out
type ActiveLocation struct {
	gorm.Model
	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"type:text"`
	Latitude    float64 `gorm:"not null"`
	Longitude   float64 `gorm:"not null"`
}

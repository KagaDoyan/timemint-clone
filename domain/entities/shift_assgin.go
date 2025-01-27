package entities

import "gorm.io/gorm"

type ShiftAssignment struct {
	gorm.Model
	Employee  []Employee `gorm:"many2many:shift_assignments;"`
	ShiftID   uint       `gorm:"not null;foreignKey:ShiftID"`
	StartDate string     `gorm:"not null"`     // For one-time shifts
	EndDate   string     `gorm:"not null"`     // For one-time shifts
	IsActive  bool       `gorm:"default:true"` // Marks whether the shift assignment is active
	Shift     Shift      `gorm:"foreignKey:ShiftID"`
}

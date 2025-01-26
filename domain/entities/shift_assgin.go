package entities

import "gorm.io/gorm"

type ShiftAssignment struct {
	gorm.Model
	EmployeeID uint   `gorm:"not null;foreignKey:EmployeeID"`
	ShiftID    uint   `gorm:"not null;foreignKey:ShiftID"`
	StartDate  string `gorm:"not null"`     // For one-time shifts
	EndDate    string `gorm:"not null"`     // For one-time shifts
	IsActive   bool   `gorm:"default:true"` // Marks whether the shift assignment is active

	Employee Employee `gorm:"foreignKey:EmployeeID"`
	Shift    Shift    `gorm:"foreignKey:ShiftID"`
}

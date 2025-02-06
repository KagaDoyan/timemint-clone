package entities

import (
	"time"

	"gorm.io/gorm"
)

type ShiftAssignment struct {
	gorm.Model
	EmployeeID    uint
	Employee      Employee
	ShiftID       uint
	Shift         Shift
	Date          time.Time
	CreatedBy     uint
	CreatedByUser Employee `gorm:"foreignKey:CreatedBy"`
}

package entities

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	EmployeeID   uint
	Employee     Employee
	Date         string `gorm:"unique"`
	CheckInTime  string
	CheckOutTime string
	Status       string
	IsLate       bool
	IsLeaveEarly bool
}

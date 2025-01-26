package entities

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	EmployeeID uint
	Employee   Employee
	CheckIn    string
	CheckOut   string
	Status     string
}

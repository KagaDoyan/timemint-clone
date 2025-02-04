package entities

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeNo string
	Dob        string
	Name       string
	Email      string
	Phone      string
	Address    string
	Password   string
	Position   string
	RoleID     uint
	Role       Role
	Department string
}

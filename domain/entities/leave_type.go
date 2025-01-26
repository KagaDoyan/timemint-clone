package entities

import "gorm.io/gorm"

type LeaveType struct {
	gorm.Model
	LeaveType   string
	Description string
	Payable     bool
}

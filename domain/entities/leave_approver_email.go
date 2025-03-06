package entities

import "gorm.io/gorm"

type LeaveApproverEmail struct {
	gorm.Model
	Email string
}

package entities

import "gorm.io/gorm"

type LeaveRequest struct {
	gorm.Model
	EmployeeID  uint `gorm:"not null;foreignKey:EmployeeID"`
	LeaveType   LeaveType
	LeaveTypeID uint     `gorm:"not null;foreignKey:LeaveTypeID"`
	StartDate   string   `gorm:"not null"`         // Start date of the leave
	EndDate     string   `gorm:"not null"`         // End date of the leave
	Status      string   `gorm:"size:20;not null"` // e.g., "Pending", "Approved", "Rejected"
	Reason      string   `gorm:"type:text"`        // Reason for leave (optional)
	Employee    Employee `gorm:"foreignKey:EmployeeID"`
}

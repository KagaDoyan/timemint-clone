package entities

import "gorm.io/gorm"

type LeaveRequest struct {
	gorm.Model
	EmployeeID  uint      // Foreign key referencing the employee who requested leave
	LeaveType   LeaveType // Relationship to the LeaveType entity
	LeaveTypeID uint      `gorm:"not null"`         // Foreign key for LeaveType
	StartDate   string    `gorm:"not null"`         // Start date of the leave
	EndDate     string    `gorm:"not null"`         // End date of the leave
	FullDay     bool      `gorm:"not null"`         // Is it a full day leave?
	Status      string    `gorm:"size:20;not null"` // Status: "Pending", "Approved", "Rejected"
	Reason      string    `gorm:"type:text"`        // Reason for leave (optional)
	Employee    Employee
	ReviewerID  *uint     // Nullable foreign key for Reviewer
	Reviewer    *Employee `gorm:"foreignKey:ReviewerID"` // Nullable reference to the reviewing employee
}

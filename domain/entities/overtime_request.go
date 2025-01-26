package entities

import "gorm.io/gorm"

type OvertimeRequest struct {
	gorm.Model
	EmployeeID    uint     `gorm:"not null;foreignKey:EmployeeID"`
	OvertimeHours float64  `gorm:"not null"`         // Number of overtime hours requested
	DateRequested string   `gorm:"not null"`         // Date the overtime is requested for
	Status        string   `gorm:"size:20;not null"` // e.g., "Pending", "Approved", "Rejected"
	Reason        string   `gorm:"type:text"`        // Reason for requesting overtime (optional)
	Employee      Employee `gorm:"foreignKey:EmployeeID"`
}

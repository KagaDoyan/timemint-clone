package models

import "time"

type LeaveRequest struct {
	ID          uint      `json:"id"`
	EmployeeID  uint      `json:"employee_id"`
	Employee    Employee  `json:"employee"`
	LeaveTypeID uint      `json:"leave_type_id"`
	LeaveType   LeaveType `json:"leave_type"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status"`
	FullDay     bool      `json:"full_day"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ReviewerID  uint      `json:"reviewer_id"`
	Reviewer    Employee  `json:"reviewer"`
}

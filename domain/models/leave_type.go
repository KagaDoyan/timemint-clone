package models

import "time"

type LeaveType struct {
	ID          uint   `json:"id"`
	LeaveType   string `json:"leave_type"`
	Description string `json:"description"`
	Payable     bool   `json:"payable"`
	AnnuallyMax int    `json:"annually_max"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

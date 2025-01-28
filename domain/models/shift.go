package models

import "time"

type Shift struct {
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	StartTime    string     `json:"start_time"`
	EndTime      string     `json:"end_time"`
	Color        string     `json:"color"`
	DepartmentID uint       `json:"department_id"`
	Department   Department `json:"department"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

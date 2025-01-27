package models

type AttendancePolicy struct {
	ID                 uint   `json:"id"`
	UpdatedAt          string `json:"updated_at"`
	MaxLateMinutes     int    `json:"max_late_minutes"`
	MinWorkHoursPerDay int    `json:"min_work_hours_per_day"`
	OvertimeThreshold  int    `json:"overtime_threshold"`
}

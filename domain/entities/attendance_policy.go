package entities

import "gorm.io/gorm"

type AttendancePolicy struct {
	gorm.Model
	MaxLateMinutes     int `json:"max_late_minutes"`
	MinWorkHoursPerDay int `json:"min_work_hours_per_day"`
	OvertimeThreshold  int `json:"overtime_threshold"`
}

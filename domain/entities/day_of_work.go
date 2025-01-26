package entities

import "gorm.io/gorm"

type DayOfWork struct {
	gorm.Model
	Day       string
	StartTime string
	EndTime   string
	IsWorkDay bool
}

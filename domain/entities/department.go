package entities

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name string `gorm:"size:100;not null"`
}

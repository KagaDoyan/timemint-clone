package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type AttendancePolicyRepository interface {
	Find() (*entities.AttendancePolicy, error)
	Update(id uint, attendancePolicy *entities.AttendancePolicy) error
}

type attendancePolicyRepository struct {
	db *gorm.DB
}

func NewAttendancePolicyRepository(db *gorm.DB) AttendancePolicyRepository {
	db.AutoMigrate(&entities.AttendancePolicy{})
	return &attendancePolicyRepository{db: db}
}

func (s attendancePolicyRepository) Find() (*entities.AttendancePolicy, error) {
	var result entities.AttendancePolicy
	err := s.db.First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a attendancePolicyRepository) Update(id uint, attendancePolicy *entities.AttendancePolicy) error {
	err := a.db.Model(&entities.AttendancePolicy{}).Select("max_late_minutes", "min_work_hours_per_day", "overtime_threshold").Where("id = ?", id).Updates(&attendancePolicy).Error
	if err != nil {
		return err
	}
	var result entities.AttendancePolicy
	err = a.db.First(&result, id).Error
	if err != nil {
		return err
	}
	return nil
}

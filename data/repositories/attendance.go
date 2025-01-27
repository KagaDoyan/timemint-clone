package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	CreateAttendance(attendance entities.Attendance) (*entities.Attendance, error)
	UpdateAttendance(id uint, attendance entities.Attendance) (*entities.Attendance, error)
	GetAttendanceByDate(date string, employeeID uint) (*entities.Attendance, error)
}

type attendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepository{db: db}
}

func (a attendanceRepository) CreateAttendance(attendance entities.Attendance) (*entities.Attendance, error) {
	err := a.db.Create(&attendance).Error
	if err != nil {
		return nil, err
	}
	var result entities.Attendance
	err = a.db.First(&result, attendance.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a attendanceRepository) UpdateAttendance(id uint, attendance entities.Attendance) (*entities.Attendance, error) {
	err := a.db.Model(&entities.Attendance{}).Where("id = ?", id).Updates(&attendance).Error
	if err != nil {
		return nil, err
	}
	var result entities.Attendance
	err = a.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (a attendanceRepository) GetAttendanceByDate(date string, employeeID uint) (*entities.Attendance, error) {
	var result entities.Attendance
	err := a.db.Where("date = ? AND employee_id = ?", date, employeeID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

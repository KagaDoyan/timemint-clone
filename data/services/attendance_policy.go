package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
)

type AttendancePolicyService interface {
	Find() (*models.AttendancePolicy, error)
	Update(id uint, attendancePolicy *models.AttendancePolicy) error
}

type attendancePolicyService struct {
	attendancePolicyRepo repositories.AttendancePolicyRepository
}

func NewAttendancePolicyService(attendancePolicyRepo repositories.AttendancePolicyRepository) AttendancePolicyService {
	return &attendancePolicyService{attendancePolicyRepo: attendancePolicyRepo}
}

func (s attendancePolicyService) Find() (*models.AttendancePolicy, error) {
	result, err := s.attendancePolicyRepo.Find()
	if err != nil {
		return nil, err
	}
	return &models.AttendancePolicy{
		ID:                 result.ID,
		MaxLateMinutes:     result.MaxLateMinutes,
		MinWorkHoursPerDay: result.MinWorkHoursPerDay,
		OvertimeThreshold:  result.OvertimeThreshold,
		UpdatedAt:          result.UpdatedAt.Format("02-01-2006 15:04:05"),
	}, nil
}

func (s attendancePolicyService) Update(id uint, attendancePolicy *models.AttendancePolicy) error {
	return s.attendancePolicyRepo.Update(id, &entities.AttendancePolicy{
		MaxLateMinutes:     attendancePolicy.MaxLateMinutes,
		MinWorkHoursPerDay: attendancePolicy.MinWorkHoursPerDay,
		OvertimeThreshold:  attendancePolicy.OvertimeThreshold,
	})
}

package services

import (
	"errors"
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"time"

	"gorm.io/gorm"
)

type AttendanceService interface {
}

type attendanceService struct {
	attendanceRepo repositories.AttendanceRepository
	workerRepo     repositories.DayOfWorkRepository
	holidayRepo    repositories.HolidayRepository
	policyRepo     repositories.AttendancePolicyRepository
}

func NewAttendanceService(attendanceRepo repositories.AttendanceRepository, workerRepo repositories.DayOfWorkRepository, holidayRepo repositories.HolidayRepository, policyRepo repositories.AttendancePolicyRepository) AttendanceService {
	return &attendanceService{
		attendanceRepo: attendanceRepo,
		workerRepo:     workerRepo,
		holidayRepo:    holidayRepo,
		policyRepo:     policyRepo,
	}
}

func (s attendanceService) CheckIn(employeeID uint) (*entities.Attendance, error) {
	today := time.Now()
	isholiday, err := s.holidayRepo.IsHoliday(today)
	if err != nil {
		return nil, err
	}
	if isholiday {
		return nil, errors.New("today is holiday go spend the day")
	}
	policy, err := s.policyRepo.Find()
	if err != nil {
		return nil, err
	}
	workday, err := s.workerRepo.FindByDay(today)
	if err != nil {
		return nil, err
	}
	attendance, err := s.attendanceRepo.GetAttendanceByDate(today.Format("02-01-2006"), employeeID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if attendance == nil {
		is_late := false
		// allow late time based on policy
		allow_late := policy.MaxLateMinutes
		// parse workday start time
		start_time, err := time.Parse("15:04", workday.StartTime)
		if err != nil {
			return nil, err
		}
		// Calculate the difference between now and start time
		time_diff := today.Sub(start_time)
		// Check if it's late based on the time difference and allowed late time
		if time_diff.Minutes() > float64(allow_late) {
			is_late = true
		}
		attendanceEntity := entities.Attendance{
			EmployeeID:  employeeID,
			Date:        today.Format("02-01-2006"),
			CheckInTime: today.Format("15:04"),
			IsLate:      is_late,
		}
		attendance, err = s.attendanceRepo.CreateAttendance(attendanceEntity)
		if err != nil {
			return nil, err
		}
		return attendance, nil
	}
	return attendance, nil
}

func (s attendanceService) CheckOut(employeeID uint) (*entities.Attendance, error) {
	today := time.Now()
	attendance, err := s.attendanceRepo.GetAttendanceByDate(today.Format("02-01-2006"), employeeID)
	if err != nil {
		return nil, err
	}
	attendance.CheckOutTime = today.Format("15:04")
	attendance.Status = "present"
	attendance, err = s.attendanceRepo.UpdateAttendance(attendance.ID, *attendance)
	if err != nil {
		return nil, err
	}
	return attendance, nil
}

package services

import (
	"fmt"
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
	"time"
)

type shiftAssignService struct {
	repository repositories.ShiftAssignRepository
}

type ShiftAssignService interface {
	Create(shiftAssign models.ShiftAssignment) (*models.ShiftAssignment, error)
	CreateBatch(created_by uint, shiftAssigns []models.ShiftAssignment) (int, int, int, []string)
	FindAll(page, limit int) ([]models.ShiftAssignment, int64, error)
	FindById(id uint) (*models.ShiftAssignment, error)
	Delete(id uint) error
	CalendarShift(month, year int) ([]models.ShiftAssignment, error)
	ShiftAssignmentReport(start, end string) ([]models.ShiftAssignment, error)
}

func NewShiftAssignService(repository repositories.ShiftAssignRepository) ShiftAssignService {
	return &shiftAssignService{repository}
}

func (s shiftAssignService) CreateBatch(craetedBy uint, shiftAssigns []models.ShiftAssignment) (int, int, int, []string) {
	failures := []string{}
	totalrecords := len(shiftAssigns)
	inserted := 0
	failed := 0
	for _, shiftAssign := range shiftAssigns {
		failstring := fmt.Sprintf("failed to insert record: [empID %v] [shiftID %v] [date %v]", shiftAssign.EmployeeID, shiftAssign.ShiftID, shiftAssign.Date)
		date, err := time.Parse("02-01-2006", shiftAssign.Date)
		if err != nil {
			failures = append(failures, failstring)
			failed++
			continue
		}
		_, err = s.repository.Create(entities.ShiftAssignment{
			EmployeeID: shiftAssign.EmployeeID,
			ShiftID:    shiftAssign.ShiftID,
			Date:       date,
		})
		if err != nil {
			failures = append(failures, failstring)
			failed++
			continue
		}
		inserted++
	}
	return totalrecords, inserted, failed, failures
}

func (s shiftAssignService) Create(shiftAssign models.ShiftAssignment) (*models.ShiftAssignment, error) {
	date, err := time.Parse("02-01-2006", shiftAssign.Date)
	if err != nil {
		return nil, err
	}
	shiftAssignEntity, err := s.repository.Create(entities.ShiftAssignment{
		EmployeeID: shiftAssign.EmployeeID,
		ShiftID:    shiftAssign.ShiftID,
		CreatedBy:  shiftAssign.CreatedBy,
		Date:       date,
	})
	if err != nil {
		return nil, err
	}
	return &models.ShiftAssignment{
		ID:         shiftAssignEntity.ID,
		EmployeeID: shiftAssignEntity.EmployeeID,
		Employee: models.Employee{
			ID:   shiftAssignEntity.Employee.ID,
			Name: shiftAssignEntity.Employee.Name,
		},
		ShiftID: shiftAssignEntity.ShiftID,
		Shift: models.Shift{
			ID:          shiftAssignEntity.Shift.ID,
			Name:        shiftAssignEntity.Shift.Name,
			Description: shiftAssignEntity.Shift.Description,
			StartTime:   shiftAssignEntity.Shift.StartTime,
			EndTime:     shiftAssignEntity.Shift.EndTime,
			Color:       shiftAssignEntity.Shift.Color,
		},
		Date: shiftAssignEntity.Date.Format("02-01-2006"),
	}, nil
}

func (s shiftAssignService) FindAll(page, limit int) ([]models.ShiftAssignment, int64, error) {
	data, count, err := s.repository.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var result []models.ShiftAssignment
	for _, shiftAssign := range data {
		result = append(result, models.ShiftAssignment{
			ID:         shiftAssign.ID,
			EmployeeID: shiftAssign.EmployeeID,
			Employee: models.Employee{
				ID:       shiftAssign.Employee.ID,
				Name:     shiftAssign.Employee.Name,
				Email:    shiftAssign.Employee.Email,
				Phone:    shiftAssign.Employee.Phone,
				Address:  shiftAssign.Employee.Address,
				Position: shiftAssign.Employee.Position,
				RoleID:   shiftAssign.Employee.RoleID,
				Role: models.Role{
					ID:   shiftAssign.Employee.Role.ID,
					Name: shiftAssign.Employee.Role.Name,
				},
				Department: shiftAssign.Employee.Department,
			},
			ShiftID: shiftAssign.ShiftID,
			Shift: models.Shift{
				ID:          shiftAssign.Shift.ID,
				Name:        shiftAssign.Shift.Name,
				Description: shiftAssign.Shift.Description,
				StartTime:   shiftAssign.Shift.StartTime,
				EndTime:     shiftAssign.Shift.EndTime,
				Color:       shiftAssign.Shift.Color,
			},
			Date: shiftAssign.Date.Format("02-01-2006"),
		})
	}
	return result, count, nil
}

func (s shiftAssignService) FindById(id uint) (*models.ShiftAssignment, error) {
	shiftAssign, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return &models.ShiftAssignment{
		ID:         shiftAssign.ID,
		EmployeeID: shiftAssign.EmployeeID,
		Employee: models.Employee{
			ID:   shiftAssign.Employee.ID,
			Name: shiftAssign.Employee.Name,
		},
		ShiftID: shiftAssign.ShiftID,
		Shift: models.Shift{
			ID:          shiftAssign.Shift.ID,
			Name:        shiftAssign.Shift.Name,
			Description: shiftAssign.Shift.Description,
			StartTime:   shiftAssign.Shift.StartTime,
			EndTime:     shiftAssign.Shift.EndTime,
			Color:       shiftAssign.Shift.Color,
		},
		Date: shiftAssign.Date.Format("02-01-2006"),
	}, nil
}

func (s shiftAssignService) Delete(id uint) error {
	return s.repository.Delete(id)
}

func (s shiftAssignService) CalendarShift(month int, year int) ([]models.ShiftAssignment, error) {
	shiftAssigns, err := s.repository.CalendarShift(month, year)
	if err != nil {
		return nil, err
	}
	var result []models.ShiftAssignment
	for _, shiftAssign := range shiftAssigns {
		result = append(result, models.ShiftAssignment{
			ID:         shiftAssign.ID,
			EmployeeID: shiftAssign.EmployeeID,
			Employee: models.Employee{
				ID:       shiftAssign.Employee.ID,
				Name:     shiftAssign.Employee.Name,
				Email:    shiftAssign.Employee.Email,
				Phone:    shiftAssign.Employee.Phone,
				Address:  shiftAssign.Employee.Address,
				Position: shiftAssign.Employee.Position,
				RoleID:   shiftAssign.Employee.RoleID,
			},
			ShiftID: shiftAssign.ShiftID,
			Shift: models.Shift{
				ID:          shiftAssign.Shift.ID,
				Name:        shiftAssign.Shift.Name,
				Description: shiftAssign.Shift.Description,
				StartTime:   shiftAssign.Shift.StartTime,
				EndTime:     shiftAssign.Shift.EndTime,
				Color:       shiftAssign.Shift.Color,
			},
			Date: shiftAssign.Date.Format("02-01-2006"),
		})
	}
	return result, nil
}

func (s shiftAssignService) ShiftAssignmentReport(start, end string) ([]models.ShiftAssignment, error) {
	shiftAssigns, err := s.repository.ShiftAssignmentReport(start, end)
	if err != nil {
		return nil, err
	}
	var result []models.ShiftAssignment
	for _, shiftAssign := range shiftAssigns {
		result = append(result, models.ShiftAssignment{
			ID:         shiftAssign.ID,
			EmployeeID: shiftAssign.EmployeeID,
			Employee: models.Employee{
				ID:       shiftAssign.Employee.ID,
				Name:     shiftAssign.Employee.Name,
				Email:    shiftAssign.Employee.Email,
				Phone:    shiftAssign.Employee.Phone,
				Address:  shiftAssign.Employee.Address,
				Position: shiftAssign.Employee.Position,
				RoleID:   shiftAssign.Employee.RoleID,
			},
			ShiftID: shiftAssign.ShiftID,
			Shift: models.Shift{
				ID:          shiftAssign.Shift.ID,
				Name:        shiftAssign.Shift.Name,
				Description: shiftAssign.Shift.Description,
				StartTime:   shiftAssign.Shift.StartTime,
				EndTime:     shiftAssign.Shift.EndTime,
				Color:       shiftAssign.Shift.Color,
			},
			Date: shiftAssign.Date.Format("02-01-2006"),
			CreatedByUser: models.Employee{
				ID:       shiftAssign.CreatedByUser.ID,
				Name:     shiftAssign.CreatedByUser.Name,
				Email:    shiftAssign.CreatedByUser.Email,
				Phone:    shiftAssign.CreatedByUser.Phone,
				Address:  shiftAssign.CreatedByUser.Address,
				Position: shiftAssign.CreatedByUser.Position,
				RoleID:   shiftAssign.CreatedByUser.RoleID,
			},
		})
	}
	return result, nil
}

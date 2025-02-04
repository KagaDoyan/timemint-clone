package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
	"time"
)

type leaveRequestService struct {
	repo repositories.LeaveRequestRepository
}

type LeaveRequestService interface {
	EmpLeaveRequests(employeeID uint, leave models.LeaveRequest) (*models.LeaveRequest, error)
	CraeteLeaveRequests(created_by uint, leave models.LeaveRequest) (*models.LeaveRequest, error)
	Update(id uint, leaveRequest models.LeaveRequest) (*models.LeaveRequest, error)
	Delete(id uint) error
	FindAll(page, limit int, status string, employeeID uint, from, to string) ([]models.LeaveRequest, int64, int64, int64, int64, error)
	FindById(id uint) (*models.LeaveRequest, error)
	CalendarLeaves(month, year int) ([]models.LeaveRequest, error)
}

func NewLeaveRequestService(repo repositories.LeaveRequestRepository) LeaveRequestService {
	return leaveRequestService{repo: repo}
}

func (s leaveRequestService) CraeteLeaveRequests(created_by uint, leave models.LeaveRequest) (*models.LeaveRequest, error) {
	leaveRequstEntity := entities.LeaveRequest{
		EmployeeID:  leave.EmployeeID,
		LeaveTypeID: leave.LeaveTypeID,
		StartDate:   leave.StartDate,
		EndDate:     leave.EndDate,
		Reason:      leave.Reason,
		Status:      "approved",
		FullDay:     leave.FullDay,
		ReviewerID:  &created_by,
	}

	result, err := s.repo.Create(leaveRequstEntity)
	if err != nil {
		return nil, err
	}

	var reviewer models.Employee
	if result.ReviewerID != nil {
		reviewer = models.Employee{
			ID:   *result.ReviewerID,
			Name: result.Reviewer.Name, // Ensure result.Reviewer is not nil
		}
	}

	return &models.LeaveRequest{
		ID:         result.ID,
		EmployeeID: result.EmployeeID,
		Employee: models.Employee{
			ID:   result.EmployeeID,
			Name: result.Employee.Name,
		},
		LeaveType: models.LeaveType{
			ID:          result.LeaveTypeID,
			LeaveType:   result.LeaveType.LeaveType,
			Description: result.LeaveType.Description,
			Payable:     result.LeaveType.Payable,
			AnnuallyMax: result.LeaveType.AnnuallyMax,
		},
		LeaveTypeID: result.LeaveTypeID,
		StartDate:   result.StartDate,
		EndDate:     result.EndDate,
		Reason:      result.Reason,
		Status:      result.Status,
		FullDay:     result.FullDay,
		Reviewer:    reviewer,
	}, nil
}

func (s leaveRequestService) EmpLeaveRequests(employeeID uint, leave models.LeaveRequest) (*models.LeaveRequest, error) {
	leaveRequstEntity := entities.LeaveRequest{
		EmployeeID:  employeeID,
		LeaveTypeID: leave.LeaveTypeID,
		StartDate:   leave.StartDate,
		EndDate:     leave.EndDate,
		Reason:      leave.Reason,
		Status:      "pending",
		FullDay:     leave.FullDay,
		ReviewerID:  nil,
	}

	result, err := s.repo.Create(leaveRequstEntity)
	if err != nil {
		return nil, err
	}

	var reviewer models.Employee
	if result.ReviewerID != nil {
		reviewer = models.Employee{
			ID:   *result.ReviewerID,
			Name: result.Reviewer.Name, // Ensure result.Reviewer is not nil
		}
	}
	return &models.LeaveRequest{
		ID:         result.ID,
		EmployeeID: result.EmployeeID,
		Employee: models.Employee{
			ID:   result.EmployeeID,
			Name: result.Employee.Name,
		},
		LeaveType: models.LeaveType{
			ID:          result.LeaveTypeID,
			LeaveType:   result.LeaveType.LeaveType,
			Description: result.LeaveType.Description,
			Payable:     result.LeaveType.Payable,
			AnnuallyMax: result.LeaveType.AnnuallyMax,
		},
		LeaveTypeID: result.LeaveTypeID,
		StartDate:   result.StartDate,
		EndDate:     result.EndDate,
		Reason:      result.Reason,
		Status:      result.Status,
		FullDay:     result.FullDay,
		Reviewer:    reviewer,
	}, nil
}

func (s leaveRequestService) Update(id uint, leaveRequest models.LeaveRequest) (*models.LeaveRequest, error) {
	leaveRequestEntity := entities.LeaveRequest{
		EmployeeID:  leaveRequest.EmployeeID,
		LeaveTypeID: leaveRequest.LeaveTypeID,
		StartDate:   leaveRequest.StartDate,
		EndDate:     leaveRequest.EndDate,
		Reason:      leaveRequest.Reason,
		Status:      leaveRequest.Status,
		Remark:      leaveRequest.Remark,
		FullDay:     leaveRequest.FullDay,
	}

	result, err := s.repo.Update(id, leaveRequestEntity)
	if err != nil {
		return nil, err
	}

	var reviewer models.Employee
	if result.Reviewer != nil {
		reviewer = models.Employee{
			ID:   *result.ReviewerID,
			Name: result.Reviewer.Name,
		}
	}
	return &models.LeaveRequest{
		ID:         result.ID,
		EmployeeID: result.EmployeeID,
		Employee: models.Employee{
			ID:   result.EmployeeID,
			Name: result.Employee.Name,
		},
		LeaveType: models.LeaveType{
			ID:          result.LeaveTypeID,
			LeaveType:   result.LeaveType.LeaveType,
			Description: result.LeaveType.Description,
			Payable:     result.LeaveType.Payable,
			AnnuallyMax: result.LeaveType.AnnuallyMax,
		},
		LeaveTypeID: result.LeaveTypeID,
		StartDate:   result.StartDate,
		EndDate:     result.EndDate,
		Reason:      result.Reason,
		Status:      result.Status,
		FullDay:     result.FullDay,
		Reviewer:    reviewer,
	}, nil
}

func (s leaveRequestService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s leaveRequestService) FindAll(page, limit int, status string, employeeID uint, from, to string) ([]models.LeaveRequest, int64, int64, int64, int64, error) {
	if len(from) > 0 && len(to) > 0 {
		fromDate, err := time.Parse("02-01-2006", from)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}
		toDate, err := time.Parse("02-01-2006", to)
		if err != nil {
			return nil, 0, 0, 0, 0, err
		}
		from = fromDate.Format("02-01-2006")
		to = toDate.Format("02-01-2006")
	}
	results, total, totalPending, totalApproved, totalRejected, err := s.repo.FindAll(page, limit, status, employeeID, from, to)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	var leaveRequests []models.LeaveRequest
	for _, result := range results {

		var reviewer models.Employee
		if result.Reviewer != nil {
			reviewer = models.Employee{
				ID:   *result.ReviewerID,
				Name: result.Reviewer.Name,
			}
		}
		leaveRequests = append(leaveRequests, models.LeaveRequest{
			ID:         result.ID,
			EmployeeID: result.EmployeeID,
			Employee: models.Employee{
				ID:         result.EmployeeID,
				EmployeeNo: result.Employee.EmployeeNo,
				Name:       result.Employee.Name,
			},
			LeaveType: models.LeaveType{
				ID:          result.LeaveTypeID,
				LeaveType:   result.LeaveType.LeaveType,
				Description: result.LeaveType.Description,
				Payable:     result.LeaveType.Payable,
				AnnuallyMax: result.LeaveType.AnnuallyMax,
			},
			LeaveTypeID: result.LeaveTypeID,
			StartDate:   result.StartDate,
			EndDate:     result.EndDate,
			Reason:      result.Reason,
			Status:      result.Status,
			FullDay:     result.FullDay,
			Reviewer:    reviewer,
		})
	}
	return leaveRequests, total, totalPending, totalApproved, totalRejected, nil
}

func (s leaveRequestService) FindById(id uint) (*models.LeaveRequest, error) {
	result, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	var reviewer models.Employee
	if result.Reviewer != nil {
		reviewer = models.Employee{
			ID:   *result.ReviewerID,
			Name: result.Reviewer.Name,
		}
	}
	return &models.LeaveRequest{
		ID:         result.ID,
		EmployeeID: result.EmployeeID,
		Employee: models.Employee{
			ID:   result.EmployeeID,
			Name: result.Employee.Name,
		},
		LeaveType: models.LeaveType{
			ID:          result.LeaveTypeID,
			LeaveType:   result.LeaveType.LeaveType,
			Description: result.LeaveType.Description,
			Payable:     result.LeaveType.Payable,
			AnnuallyMax: result.LeaveType.AnnuallyMax,
		},
		LeaveTypeID: result.LeaveTypeID,
		StartDate:   result.StartDate,
		EndDate:     result.EndDate,
		Reason:      result.Reason,
		Status:      result.Status,
		FullDay:     result.FullDay,
		Reviewer:    reviewer,
	}, nil
}

func (s leaveRequestService) CalendarLeaves(month, year int) ([]models.LeaveRequest, error) {
	datas, err := s.repo.CalendarLeaves(month, year)
	if err != nil {
		return nil, err
	}
	var results []models.LeaveRequest
	for _, data := range datas {
		results = append(results, models.LeaveRequest{
			ID:         data.ID,
			EmployeeID: data.EmployeeID,
			Employee: models.Employee{
				ID:   data.EmployeeID,
				Name: data.Employee.Name,
			},
			LeaveType: models.LeaveType{
				ID:          data.LeaveTypeID,
				LeaveType:   data.LeaveType.LeaveType,
				Description: data.LeaveType.Description,
				Payable:     data.LeaveType.Payable,
				AnnuallyMax: data.LeaveType.AnnuallyMax,
			},
			LeaveTypeID: data.LeaveTypeID,
			StartDate:   data.StartDate,
			EndDate:     data.EndDate,
			Reason:      data.Reason,
			Status:      data.Status,
			FullDay:     data.FullDay,
		})
	}
	return results, nil
}

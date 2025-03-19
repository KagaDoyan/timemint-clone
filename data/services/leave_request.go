package services

import (
	"fmt"
	"go-fiber/bootstrap"
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
	"time"
)

type leaveRequestService struct {
	repo repositories.LeaveRequestRepository
}

var leave_notification = bootstrap.GetEnv("leave_notification", "")

type LeaveRequestService interface {
	EmpLeaveRequests(employeeID uint, leave models.LeaveRequest) (*models.LeaveRequest, error)
	CraeteLeaveRequests(created_by uint, leave models.LeaveRequest) (*models.LeaveRequest, error)
	Update(id uint, leaveRequest models.LeaveRequest) (*models.LeaveRequest, error)
	Delete(id uint) error
	FindAll(page, limit int, status string, employeeID uint, from, to string) ([]models.LeaveRequest, int64, int64, int64, int64, error)
	FindById(id uint) (*models.LeaveRequest, error)
	CalendarLeaves(month, year int) ([]models.LeaveRequest, error)
	LeaveRequestReport(start, end string) ([]models.LeaveRequest, error)
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

	approvers, _ := s.repo.ApproverEmails()
	if len(approvers) > 0 {
		//send email
		email_body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Leave Request Notification</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					line-height: 1.6;
					margin: 0;
					padding: 0;
					background-color: #f9f9f9;
				}
				.email-container {
					max-width: 600px;
					margin: 20px auto;
					padding: 20px;
					background-color: #ffffff;
					border: 1px solid #dddddd;
					border-radius: 8px;
				}
				.header {
					background-color: #4CAF50;
					color: #ffffff;
					padding: 10px;
					text-align: center;
					border-radius: 8px 8px 0 0;
				}
				.content {
					padding: 20px;
				}
				.footer {
					font-size: 0.9em;
					color: #888888;
					text-align: center;
					margin-top: 20px;
				}
				.button {
					display: inline-block;
					padding: 10px 20px;
					margin-top: 20px;
					background-color: #4CAF50;
					color: #ffffff;
					text-decoration: none;
					border-radius: 5px;
				}
				.button:hover {
					background-color: #45a049;
				}
			</style>
		</head>
		<body>
			<div class="email-container">
				<div class="header">
					<h1>Leave Request Notification</h1>
				</div>
				<div class="content">
					<p>You have received a new leave request from <strong>%s</strong>.</p>
					<p><strong>Details:</strong></p>
					<ul>
						<li><strong>Type of Leave:</strong> %s</li>
						<li><strong>Start Date:</strong> %s</li>
						<li><strong>End Date:</strong> %s</li>
						<li><strong>Reason:</strong> %s</li>
					</ul>
					<p>Please review the request and take necessary action.</p>
					<a href="https://attendance.homekvs.pw/leave-approval" class="button">Go to Approval Page</a>
					<p>Thank you!</p>
				</div>
				<div class="footer">
					<p>This is an automated message. Please do not reply directly to this email.</p>
				</div>
			</div>
		</body>
		</html>
		`, result.Employee.Name, result.LeaveType.LeaveType, result.StartDate, result.EndDate, result.Reason)
		for _, approver := range approvers {
			go SendEmail(email_body, approver.Email)
		}
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
		ReviewerID:  &leaveRequest.ReviewerID,
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

func (s leaveRequestService) LeaveRequestReport(start, end string) ([]models.LeaveRequest, error) {
	datas, err := s.repo.LeaveRequestReport(start, end)
	if err != nil {
		return nil, err
	}
	var results []models.LeaveRequest
	for _, data := range datas {
		var reviewer models.Employee
		if data.Reviewer != nil {
			reviewer = models.Employee{
				ID:   *data.ReviewerID,
				Name: data.Reviewer.Name,
			}
		}
		results = append(results, models.LeaveRequest{
			ID:         data.ID,
			EmployeeID: data.EmployeeID,
			Employee: models.Employee{
				EmployeeNo: data.Employee.EmployeeNo,
				ID:         data.EmployeeID,
				Name:       data.Employee.Name,
				Email:      data.Employee.Email,
				Phone:      data.Employee.Phone,
				Address:    data.Employee.Address,
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
			Reviewer:    reviewer,
		})
	}
	return results, nil
}

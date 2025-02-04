package repositories

import (
	"go-fiber/domain/entities"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type leaveRequestRepository struct {
	db *gorm.DB
}

type LeaveRequestRepository interface {
	Create(leaveRequest entities.LeaveRequest) (*entities.LeaveRequest, error)
	Update(id uint, leaveRequest entities.LeaveRequest) (*entities.LeaveRequest, error)
	Delete(id uint) error
	FindAll(page, limit int, status string, employeeID uint, from, to string) ([]entities.LeaveRequest, int64, int64, int64, int64, error)
	FindById(id uint) (*entities.LeaveRequest, error)
	CalendarLeaves(month, year int) ([]entities.LeaveRequest, error)
}

func NewLeaveRequestRepository(db *gorm.DB) LeaveRequestRepository {
	db.AutoMigrate(&entities.LeaveRequest{})
	return &leaveRequestRepository{db: db}
}

func (r leaveRequestRepository) Create(leaveRequest entities.LeaveRequest) (*entities.LeaveRequest, error) {
	err := r.db.Create(&leaveRequest).Error
	if err != nil {
		return nil, err
	}
	var result entities.LeaveRequest
	err = r.db.Preload(clause.Associations).First(&result, leaveRequest.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r leaveRequestRepository) Update(id uint, leaveRequest entities.LeaveRequest) (*entities.LeaveRequest, error) {
	err := r.db.Model(&entities.LeaveRequest{}).Where("id = ?", id).Select("status", "reviewer_id", "remark").Updates(&leaveRequest).Error
	if err != nil {
		return nil, err
	}
	var result entities.LeaveRequest
	err = r.db.Preload(clause.Associations).First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r leaveRequestRepository) Delete(id uint) error {
	err := r.db.Delete(&entities.LeaveRequest{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r leaveRequestRepository) FindAll(page, limit int, status string, employeeID uint, from, to string) ([]entities.LeaveRequest, int64, int64, int64, int64, error) {
	var statusWhere string
	var employeeIDWhere string
	var fromDateWhere string
	var toDateWhere string
	var totalPending int64
	var totalApproved int64
	var totalRejected int64

	if len(from) > 0 && len(to) > 0 {
		fromDateWhere = "STR_TO_DATE(start_date, '%d-%m-%Y') >= STR_TO_DATE('" + from + "', '%d-%m-%Y')"
		toDateWhere = "STR_TO_DATE(start_date, '%d-%m-%Y') <= STR_TO_DATE('" + to + "', '%d-%m-%Y')"
	}
	if len(status) > 0 {
		statusWhere = "status = '" + status + "'"
	}
	if employeeID > 0 {
		employeeIDWhere = "employee_id = " + strconv.Itoa(int(employeeID))
	}

	var leaveRequests []entities.LeaveRequest
	var total int64
	err := r.db.Model(&entities.LeaveRequest{}).Where(statusWhere).Where(employeeIDWhere).Where(fromDateWhere).Where(toDateWhere).Count(&total).Error
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	err = r.db.Model(&entities.LeaveRequest{}).Where("status = 'pending'").Where(employeeIDWhere).Where(fromDateWhere).Where(toDateWhere).Count(&totalPending).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			totalPending = 0
		} else {
			return nil, 0, 0, 0, 0, err
		}
	}
	err = r.db.Model(&entities.LeaveRequest{}).Where("status = 'approved'").Where(employeeIDWhere).Where(fromDateWhere).Where(toDateWhere).Count(&totalApproved).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			totalApproved = 0
		} else {
			return nil, 0, 0, 0, 0, err
		}
	}
	err = r.db.Model(&entities.LeaveRequest{}).Where("status = 'rejected'").Where(employeeIDWhere).Where(fromDateWhere).Where(toDateWhere).Count(&totalRejected).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			totalRejected = 0
		} else {
			return nil, 0, 0, 0, 0, err
		}
	}
	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Order("created_at desc").Preload(clause.Associations).Preload("Employee.Role").Where(statusWhere).Where(employeeIDWhere).Where(fromDateWhere).Where(toDateWhere).Find(&leaveRequests).Error
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	return leaveRequests, total, totalPending, totalApproved, totalRejected, nil
}

func (r leaveRequestRepository) FindById(id uint) (*entities.LeaveRequest, error) {
	var leaveRequest entities.LeaveRequest
	err := r.db.Preload(clause.Associations).First(&leaveRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &leaveRequest, nil
}

func (r leaveRequestRepository) CalendarLeaves(month, year int) ([]entities.LeaveRequest, error) {
	var leaveRequests []entities.LeaveRequest
	err := r.db.Preload(clause.Associations).Where("MONTH(STR_TO_DATE(start_date, '%d-%m-%Y')) = ? AND YEAR(STR_TO_DATE(start_date, '%d-%m-%Y')) = ?", month, year).
		Find(&leaveRequests).Error
	if err != nil {
		return nil, err
	}
	return leaveRequests, nil
}

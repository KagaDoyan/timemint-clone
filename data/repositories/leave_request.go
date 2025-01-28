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
	FindAll(page, limit int, status string, employeeID uint) ([]entities.LeaveRequest, int64, error)
	FindById(id uint) (*entities.LeaveRequest, error)
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
	err := r.db.Model(&entities.LeaveRequest{}).Where("id = ?", id).Select("start_date", "end_date", "reason", "status", "reviewer_id").Updates(&leaveRequest).Error
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

func (r leaveRequestRepository) FindAll(page, limit int, status string, employeeID uint) ([]entities.LeaveRequest, int64, error) {
	var statusWhere string
	var employeeIDWhere string
	if len(status) > 0 {
		statusWhere = "status = '" + status + "'"
	}
	if employeeID > 0 {
		employeeIDWhere = "employee_id = " + strconv.Itoa(int(employeeID))
	}
	var leaveRequests []entities.LeaveRequest
	var total int64
	err := r.db.Model(&entities.LeaveRequest{}).Where(statusWhere).Where(employeeIDWhere).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Preload(clause.Associations).Where(statusWhere).Where(employeeIDWhere).Find(&leaveRequests).Error
	if err != nil {
		return nil, 0, err
	}
	return leaveRequests, total, nil
}

func (r leaveRequestRepository) FindById(id uint) (*entities.LeaveRequest, error) {
	var leaveRequest entities.LeaveRequest
	err := r.db.Preload(clause.Associations).First(&leaveRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &leaveRequest, nil
}

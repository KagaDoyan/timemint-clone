package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type leaveTypeRepository struct {
	db *gorm.DB
}

type LeaveTypeRepository interface {
	Create(leaveType entities.LeaveType) (*entities.LeaveType, error)
	Update(id uint, leaveType entities.LeaveType) (*entities.LeaveType, error)
	Delete(id uint) error
	FindById(id uint) (*entities.LeaveType, error)
	FindAll(page, limit int) ([]entities.LeaveType, int64, error)
}

func NewLeaveTypeRepository(db *gorm.DB) LeaveTypeRepository {
	db.AutoMigrate(&entities.LeaveType{})
	return &leaveTypeRepository{db: db}
}

func (r leaveTypeRepository) Create(leaveType entities.LeaveType) (*entities.LeaveType, error) {
	err := r.db.Create(&leaveType).Error
	if err != nil {
		return nil, err
	}
	var result entities.LeaveType
	err = r.db.First(&result, leaveType.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r leaveTypeRepository) Update(id uint, leaveType entities.LeaveType) (*entities.LeaveType, error) {
	err := r.db.Model(&entities.LeaveType{}).Where("id = ?", id).Select("leave_type", "description", "payable", "annually_max").Updates(&leaveType).Error
	if err != nil {
		return nil, err
	}
	var result entities.LeaveType
	err = r.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r leaveTypeRepository) Delete(id uint) error {
	err := r.db.Unscoped().Delete(&entities.LeaveType{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r leaveTypeRepository) FindById(id uint) (*entities.LeaveType, error) {
	var leaveType entities.LeaveType
	err := r.db.First(&leaveType, id).Error
	if err != nil {
		return nil, err
	}
	return &leaveType, nil
}

func (r leaveTypeRepository) FindAll(page, limit int) ([]entities.LeaveType, int64, error) {
	var leaveTypes []entities.LeaveType
	var total int64
	err := r.db.Model(&entities.LeaveType{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&leaveTypes).Error
	if err != nil {
		return nil, 0, err
	}
	return leaveTypes, total, nil
}

package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
)

type departmentRepository struct {
	db *gorm.DB
}

type DepartmentRepository interface {
	FindAll(page, limit int) ([]entities.Department, int64, error)
	FindById(id uint) (entities.Department, error)
	Create(department entities.Department) (*entities.Department, error)
	Update(id uint, department entities.Department) (*entities.Department, error)
	Delete(id uint) error
	FindByEmployee(id uint) ([]entities.Department, error)
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	db.AutoMigrate(&entities.Department{})
	return &departmentRepository{db: db}
}

func (r departmentRepository) FindAll(page, limit int) ([]entities.Department, int64, error) {
	var departments []entities.Department
	var total int64
	err := r.db.Model(&entities.Department{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Paginate the results
	offset := (page - 1) * limit
	err = r.db.Offset(offset).Limit(limit).Find(&departments).Error
	if err != nil {
		return nil, 0, err
	}
	return departments, total, nil
}

func (r departmentRepository) FindById(id uint) (entities.Department, error) {
	var department entities.Department
	err := r.db.First(&department, id).Error
	if err != nil {
		return department, err
	}
	return department, nil
}

func (r departmentRepository) Create(department entities.Department) (*entities.Department, error) {
	err := r.db.Create(&department).Error
	if err != nil {
		return nil, err
	}
	var result entities.Department
	err = r.db.First(&result, department.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r departmentRepository) Update(id uint, department entities.Department) (*entities.Department, error) {
	err := r.db.Model(&entities.Department{}).Where("id = ?", id).Updates(&department).Error
	if err != nil {
		return nil, err
	}
	var result entities.Department
	err = r.db.First(&result, id).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r departmentRepository) Delete(id uint) error {
	err := r.db.Unscoped().Delete(&entities.Department{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r departmentRepository) FindByEmployee(id uint) ([]entities.Department, error) {
	var departments []entities.Department
	err := r.db.Where("id IN (SELECT department_id FROM employees WHERE id = ?)", id).Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

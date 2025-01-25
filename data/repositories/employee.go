package repositories

import (
	"go-fiber/domain/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type employeeRepository struct {
	db *gorm.DB
}

type EmployeeRepository interface {
	FindAll(page, limit int) ([]entities.Employee, int64, error)
	FindByID(id uint) (*entities.Employee, error)
	Create(employee *entities.Employee) (*entities.Employee, error)
	Update(employee *entities.Employee) (*entities.Employee, error)
	Delete(id uint) error
	FindByEmail(email string) (*entities.Employee, error)
}

func (s employeeRepository) FindAll(page, limit int) ([]entities.Employee, int64, error) {
	var employees []entities.Employee
	var total int64

	// Calculate offset
	offset := (page - 1) * limit

	// Count total records
	if err := s.db.Model(&entities.Employee{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Find paginated records
	err := s.db.Offset(offset).Limit(limit).Preload("Role").Find(&employees).Error
	if err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

func (s employeeRepository) FindByID(id uint) (*entities.Employee, error) {
	var employee entities.Employee
	err := s.db.Preload("Role").First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (s employeeRepository) Create(employee *entities.Employee) (*entities.Employee, error) {
	err := s.db.Create(employee).Error
	if err != nil {
		return nil, err
	}
	var result entities.Employee
	err = s.db.Preload("Role").First(&result, employee.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s employeeRepository) Update(employee *entities.Employee) (*entities.Employee, error) {
	//use update function
	err := s.db.Select("name", "email", "phone", "address", "position", "role_id", "employee_id").Updates(employee).Error
	if err != nil {
		return nil, err
	}
	var result entities.Employee
	err = s.db.Preload("Role").First(&result, employee.ID).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s employeeRepository) Delete(id uint) error {
	//hard delete
	err := s.db.Unscoped().Delete(&entities.Employee{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (s employeeRepository) FindByEmail(email string) (*entities.Employee, error) {
	var employee entities.Employee
	err := s.db.Where("email = ?", email).Preload(clause.Associations).First(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	db.AutoMigrate(&entities.Employee{})
	return &employeeRepository{db: db}
}

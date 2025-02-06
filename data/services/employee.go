package services

import (
	"errors"
	"fmt"
	"go-fiber/core/logs"
	"go-fiber/core/utilities"
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"

	"gorm.io/gorm"
)

type EmployeeService interface {
	FindAll(page, limit int) ([]models.Employee, int64, error)
	FindByID(id uint) (*models.Employee, error)
	Create(employee *models.Employee) (*models.Employee, error)
	Update(employee *models.Employee) (*models.Employee, error)
	Delete(id uint) error
	FindByEmail(email string) (*models.Employee, error)
	Login(email, password string) (*models.Employee, error)
	Option() ([]models.Employee, error)
	EmployeeReport() ([]models.Employee, error)
	SetPassword(email, password string) (*models.Employee, error)
}

type employeeService struct {
	employeerepo repositories.EmployeeRepository
}

func (s employeeService) Option() ([]models.Employee, error) {
	employees, err := s.employeerepo.Option()
	if err != nil {
		return nil, err
	}
	var result []models.Employee
	for _, employee := range employees {
		result = append(result, models.Employee{
			ID:         employee.ID,
			EmployeeNo: employee.EmployeeNo,
			Name:       employee.Name,
			Email:      employee.Email,
			Phone:      employee.Phone,
			Address:    employee.Address,
			Position:   employee.Position,
			RoleID:     employee.RoleID,
			Role: models.Role{
				ID:   employee.Role.ID,
				Name: employee.Role.Name,
			},
			Department: employee.Department,
		})
	}
	return result, nil
}

func (s employeeService) FindAll(page, limit int) ([]models.Employee, int64, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // Default limit
	}

	// Find employees with pagination
	employees, total, err := s.employeerepo.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	// Convert entities to models
	var result []models.Employee
	for _, employee := range employees {
		result = append(result, models.Employee{
			ID:         employee.ID,
			EmployeeNo: employee.EmployeeNo,
			Name:       employee.Name,
			Email:      employee.Email,
			Phone:      employee.Phone,
			Address:    employee.Address,
			Position:   employee.Position,
			RoleID:     employee.RoleID,
			Role: models.Role{
				ID:   employee.Role.ID,
				Name: employee.Role.Name,
			},
			Department: employee.Department,
			CreatedAt:  employee.CreatedAt,
			UpdatedAt:  employee.UpdatedAt,
		})
	}

	return result, total, nil
}

func (s employeeService) FindByID(id uint) (*models.Employee, error) {
	employee, err := s.employeerepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &models.Employee{
		ID:         employee.ID,
		EmployeeNo: employee.EmployeeNo,
		Name:       employee.Name,
		Email:      employee.Email,
		Phone:      employee.Phone,
		Address:    employee.Address,
		Position:   employee.Position,
		RoleID:     employee.RoleID,
		Role: models.Role{
			ID:   employee.Role.ID,
			Name: employee.Role.Name,
		},
		Department: employee.Department,
		CreatedAt:  employee.CreatedAt,
		UpdatedAt:  employee.UpdatedAt,
	}, nil
}

func (s employeeService) Create(employee *models.Employee) (*models.Employee, error) {
	// Convert model to entity
	entityEmployee := &entities.Employee{
		EmployeeNo: employee.EmployeeNo,
		Department: employee.Department,
		Name:       employee.Name,
		Email:      employee.Email,
		Phone:      employee.Phone,
		Address:    employee.Address,
		Position:   employee.Position,
		RoleID:     employee.RoleID,
	}
	data, err := s.employeerepo.Create(entityEmployee)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &models.Employee{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Phone:    data.Phone,
		Address:  data.Address,
		Position: data.Position,
		Role: models.Role{
			ID:   data.Role.ID,
			Name: data.Role.Name,
		},
		Department: data.Department,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}, nil
}

func (s employeeService) Update(employee *models.Employee) (*models.Employee, error) {
	// Convert model to entity
	entityEmployee := &entities.Employee{
		Model:      gorm.Model{ID: employee.ID},
		Name:       employee.Name,
		Email:      employee.Email,
		Phone:      employee.Phone,
		Address:    employee.Address,
		Position:   employee.Position,
		RoleID:     employee.RoleID,
		EmployeeNo: employee.EmployeeNo,
		Department: employee.Department,
	}
	data, err := s.employeerepo.Update(entityEmployee)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &models.Employee{
		ID:         data.ID,
		EmployeeNo: data.EmployeeNo,
		Name:       data.Name,
		Email:      data.Email,
		Phone:      data.Phone,
		Address:    data.Address,
		Position:   data.Position,
		RoleID:     data.RoleID,
		Role: models.Role{
			ID:   data.Role.ID,
			Name: data.Role.Name,
		},
		Department: data.Department,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}, nil
}

func (s employeeService) Delete(id uint) error {
	return s.employeerepo.Delete(id)
}

func (s employeeService) FindByEmail(email string) (*models.Employee, error) {
	employee, err := s.employeerepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return &models.Employee{
		ID:       employee.ID,
		Name:     employee.Name,
		Email:    employee.Email,
		Phone:    employee.Phone,
		Address:  employee.Address,
		Position: employee.Position,
		Role: models.Role{
			ID:   employee.Role.ID,
			Name: employee.Role.Name,
		},
		Department: employee.Department,
		CreatedAt:  employee.CreatedAt,
		UpdatedAt:  employee.UpdatedAt,
	}, nil
}

func (s employeeService) Login(email, password string) (*models.Employee, error) {
	// Find employee by email
	employee, err := s.employeerepo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	if len(employee.Password) == 0 {
		return nil, fmt.Errorf("set password")
	}
	// Verify password
	encodingPassword, err := utilities.GetAESEncrypted(password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if employee.Password != encodingPassword {
		return nil, fmt.Errorf("invalid credentials")
	}
	// Convert entity to model
	return &models.Employee{
		ID:       employee.ID,
		Name:     employee.Name,
		Email:    employee.Email,
		Phone:    employee.Phone,
		Address:  employee.Address,
		Position: employee.Position,
		Role: models.Role{
			ID:   employee.Role.ID,
			Name: employee.Role.Name,
		},
		Department: employee.Department,
		CreatedAt:  employee.CreatedAt,
		UpdatedAt:  employee.UpdatedAt,
	}, nil
}

func (s employeeService) EmployeeReport() ([]models.Employee, error) {
	employees, err := s.employeerepo.EmployeeReport()
	if err != nil {
		return nil, err
	}
	var result []models.Employee
	for _, employee := range employees {
		result = append(result, models.Employee{
			ID:         employee.ID,
			EmployeeNo: employee.EmployeeNo,
			Name:       employee.Name,
			Email:      employee.Email,
			Phone:      employee.Phone,
			Address:    employee.Address,
			Position:   employee.Position,
			RoleID:     employee.RoleID,
			Role: models.Role{
				ID:   employee.Role.ID,
				Name: employee.Role.Name,
			},
			Department: employee.Department,
			CreatedAt:  employee.CreatedAt,
			UpdatedAt:  employee.UpdatedAt,
		})
	}
	return result, nil
}

func (s employeeService) SetPassword(email, password string) (*models.Employee, error) {
	encodepassword, err := utilities.GetAESEncrypted(password)
	if err != nil {
		return nil, err
	}
	entityEmployee := entities.Employee{
		Password: encodepassword,
	}
	emp, err := s.employeerepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if len(emp.Password) > 0 {
		return nil, errors.New("password already set")
	}
	employee, err := s.employeerepo.SetPassword(emp.ID, entityEmployee)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &models.Employee{
		ID:         employee.ID,
		EmployeeNo: employee.EmployeeNo,
		Name:       employee.Name,
		Email:      employee.Email,
		Phone:      employee.Phone,
		Address:    employee.Address,
		Position:   employee.Position,
		RoleID:     employee.RoleID,
		Role: models.Role{
			ID:   employee.Role.ID,
			Name: employee.Role.Name,
		},
		Department: employee.Department,
		CreatedAt:  employee.CreatedAt,
		UpdatedAt:  employee.UpdatedAt,
	}, nil
}

func NewEmployeeServices(employeerepo repositories.EmployeeRepository) EmployeeService {
	return &employeeService{employeerepo: employeerepo}
}

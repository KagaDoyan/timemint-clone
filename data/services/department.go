package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
)

type departmentService struct {
	repository repositories.DepartmentRepository
}

type DepartmentService interface {
	FindAll(page, limit int) ([]models.Department, int64, error)
	FindById(id uint) (models.Department, error)
	Create(department models.Department) (*models.Department, error)
	Update(id uint, department models.Department) (*models.Department, error)
	Delete(id uint) error
}

func NewDepartmentService(repository repositories.DepartmentRepository) DepartmentService {
	return &departmentService{repository}
}

func (s departmentService) FindAll(page, limit int) ([]models.Department, int64, error) {
	departments, total, err := s.repository.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var result []models.Department
	for _, department := range departments {
		result = append(result, models.Department{
			ID:        department.ID,
			Name:      department.Name,
			CreatedAt: department.CreatedAt,
			UpdatedAt: department.UpdatedAt,
		})
	}
	return result, total, nil
}

func (s departmentService) FindById(id uint) (models.Department, error) {
	department, err := s.repository.FindById(id)
	if err != nil {
		return models.Department{}, err
	}
	return models.Department{
		ID:        department.ID,
		Name:      department.Name,
		CreatedAt: department.CreatedAt,
		UpdatedAt: department.UpdatedAt,
	}, nil
}

func (s departmentService) Create(department models.Department) (*models.Department, error) {
	entityDepartment := entities.Department{
		Name: department.Name,
	}
	result, err := s.repository.Create(entityDepartment)
	if err != nil {
		return nil, err
	}
	return &models.Department{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (s departmentService) Update(id uint, department models.Department) (*models.Department, error) {
	entityDepartment := entities.Department{
		Name: department.Name,
	}
	result, err := s.repository.Update(id, entityDepartment)
	if err != nil {
		return nil, err
	}
	return &models.Department{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (s departmentService) Delete(id uint) error {
	return s.repository.Delete(id)
}

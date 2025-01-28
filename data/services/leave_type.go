package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
)

type LeaveTypeService interface {
	Create(leaveType models.LeaveType) (*models.LeaveType, error)
	Update(id uint, leaveType models.LeaveType) (*models.LeaveType, error)
	Delete(id uint) error
	FindById(id uint) (*models.LeaveType, error)
	FindAll(page, limit int) ([]models.LeaveType, int64, error)
}

type leaveTypeService struct {
	repository repositories.LeaveTypeRepository
}

func NewLeaveTypeService(repository repositories.LeaveTypeRepository) LeaveTypeService {
	return &leaveTypeService{repository}
}

func (s leaveTypeService) Create(leaveType models.LeaveType) (*models.LeaveType, error) {
	leaveTypeEntity := entities.LeaveType{
		LeaveType:   leaveType.LeaveType,
		Description: leaveType.Description,
		Payable:     leaveType.Payable,
		AnnuallyMax: leaveType.AnnuallyMax,
	}

	result, err := s.repository.Create(leaveTypeEntity)
	if err != nil {
		return nil, err
	}
	return &models.LeaveType{
		ID:          result.ID,
		LeaveType:   result.LeaveType,
		Description: result.Description,
		Payable:     result.Payable,
		AnnuallyMax: result.AnnuallyMax,
	}, nil
}

func (s leaveTypeService) Update(id uint, leaveType models.LeaveType) (*models.LeaveType, error) {
	leaveTypeEntity := entities.LeaveType{
		LeaveType:   leaveType.LeaveType,
		Description: leaveType.Description,
		Payable:     leaveType.Payable,
		AnnuallyMax: leaveType.AnnuallyMax,
	}

	result, err := s.repository.Update(id, leaveTypeEntity)
	if err != nil {
		return nil, err
	}
	return &models.LeaveType{
		ID:          result.ID,
		LeaveType:   result.LeaveType,
		Description: result.Description,
		Payable:     result.Payable,
		AnnuallyMax: result.AnnuallyMax,
	}, nil
}

func (s leaveTypeService) Delete(id uint) error {
	return s.repository.Delete(id)
}

func (s leaveTypeService) FindById(id uint) (*models.LeaveType, error) {
	result, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return &models.LeaveType{
		ID:          result.ID,
		LeaveType:   result.LeaveType,
		Description: result.Description,
		Payable:     result.Payable,
		AnnuallyMax: result.AnnuallyMax,
	}, nil
}

func (s leaveTypeService) FindAll(page, limit int) ([]models.LeaveType, int64, error) {
	results, total, err := s.repository.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var leaveTypes []models.LeaveType
	for _, result := range results {
		leaveTypes = append(leaveTypes, models.LeaveType{
			ID:          result.ID,
			LeaveType:   result.LeaveType,
			Description: result.Description,
			Payable:     result.Payable,
			AnnuallyMax: result.AnnuallyMax,
		})
	}
	return leaveTypes, total, nil
}

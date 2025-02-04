package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
)

type ShiftService interface {
	FindAll(page, limit int) ([]models.Shift, int64, error)
	Create(shift models.Shift) (*models.Shift, error)
	FindById(id uint) (*models.Shift, error)
	Update(id uint, shift models.Shift) (*models.Shift, error)
	Delete(id uint) error
	Option() ([]models.Shift, error)
}

type shiftService struct {
	repository repositories.ShiftRepository
}

func NewShiftService(repository repositories.ShiftRepository) ShiftService {
	return &shiftService{repository}
}

func (s shiftService) Option() ([]models.Shift, error) {
	shifts, err := s.repository.Option()
	if err != nil {
		return nil, err
	}
	var result []models.Shift
	for _, shift := range shifts {
		result = append(result, models.Shift{
			ID:          shift.ID,
			Name:        shift.Name,
			Description: shift.Description,
			StartTime:   shift.StartTime,
			EndTime:     shift.EndTime,
			Color:       shift.Color,
		})
	}
	return result, nil
}

func (s shiftService) FindAll(page, limit int) ([]models.Shift, int64, error) {
	shifts, total, err := s.repository.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var result []models.Shift
	for _, shift := range shifts {
		result = append(result, models.Shift{
			ID:          shift.ID,
			Name:        shift.Name,
			Description: shift.Description,
			StartTime:   shift.StartTime,
			EndTime:     shift.EndTime,
			Color:       shift.Color,
			CreatedAt:   shift.CreatedAt,
			UpdatedAt:   shift.UpdatedAt,
		})
	}
	return result, total, nil
}

func (s shiftService) FindById(id uint) (*models.Shift, error) {
	shift, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return &models.Shift{
		ID:          shift.ID,
		Name:        shift.Name,
		Description: shift.Description,
		StartTime:   shift.StartTime,
		EndTime:     shift.EndTime,
		Color:       shift.Color,
		CreatedAt:   shift.CreatedAt,
		UpdatedAt:   shift.UpdatedAt,
	}, nil
}

func (s shiftService) Create(shift models.Shift) (*models.Shift, error) {
	entityShift := entities.Shift{
		Name:        shift.Name,
		Description: shift.Description,
		StartTime:   shift.StartTime,
		EndTime:     shift.EndTime,
		Color:       shift.Color,
	}
	result, err := s.repository.Create(entityShift)
	if err != nil {
		return nil, err
	}
	return &models.Shift{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		StartTime:   result.StartTime,
		EndTime:     result.EndTime,
		Color:       result.Color,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (s shiftService) Update(id uint, shift models.Shift) (*models.Shift, error) {
	entityShift := entities.Shift{
		Name:        shift.Name,
		Description: shift.Description,
		StartTime:   shift.StartTime,
		EndTime:     shift.EndTime,
		Color:       shift.Color,
	}
	result, err := s.repository.Update(id, entityShift)
	if err != nil {
		return nil, err
	}
	return &models.Shift{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		StartTime:   result.StartTime,
		EndTime:     result.EndTime,
		Color:       result.Color,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (s shiftService) Delete(id uint) error {
	return s.repository.Delete(id)
}

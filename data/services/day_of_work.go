package services

import (
	"go-fiber/data/repositories"
	"go-fiber/domain/entities"
	"go-fiber/domain/models"
)

type DayOfWorkService interface {
	FindAll(page, limit int) ([]models.DayOfWork, int64, error)
	FindByID(id uint) (*models.DayOfWork, error)
	Create(dayOfWork models.DayOfWork) (*models.DayOfWork, error)
	Update(id uint, dayOfWork models.DayOfWork) (*models.DayOfWork, error)
	Delete(id uint) error
}

type dayOfWorkService struct {
	repository repositories.DayOfWorkRepository
}

func NewDayOfWorkService(repository repositories.DayOfWorkRepository) DayOfWorkService {
	return &dayOfWorkService{repository}
}

func (s dayOfWorkService) FindAll(page, limit int) ([]models.DayOfWork, int64, error) {
	dayOfWorks, total, err := s.repository.FindAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	var result []models.DayOfWork
	for _, dayOfWork := range dayOfWorks {
		result = append(result, models.DayOfWork{
			ID:        dayOfWork.ID,
			Day:       dayOfWork.Day,
			StartTime: dayOfWork.StartTime,
			EndTime:   dayOfWork.EndTime,
			IsWorkDay: dayOfWork.IsWorkDay,
		})
	}
	return result, total, nil
}

func (s dayOfWorkService) FindByID(id uint) (*models.DayOfWork, error) {
	dayOfWork, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &models.DayOfWork{
		ID:        dayOfWork.ID,
		Day:       dayOfWork.Day,
		StartTime: dayOfWork.StartTime,
		EndTime:   dayOfWork.EndTime,
		IsWorkDay: dayOfWork.IsWorkDay,
	}, nil
}

func (s dayOfWorkService) Create(dayOfWork models.DayOfWork) (*models.DayOfWork, error) {
	dayOfWorkEntity := entities.DayOfWork{
		Day:       dayOfWork.Day,
		StartTime: dayOfWork.StartTime,
		EndTime:   dayOfWork.EndTime,
		IsWorkDay: dayOfWork.IsWorkDay,
	}
	result, err := s.repository.Create(dayOfWorkEntity)
	if err != nil {
		return nil, err
	}
	return &models.DayOfWork{
		ID:        result.ID,
		Day:       result.Day,
		StartTime: result.StartTime,
		EndTime:   result.EndTime,
		IsWorkDay: result.IsWorkDay,
	}, nil
}

func (s dayOfWorkService) Update(id uint, dayOfWork models.DayOfWork) (*models.DayOfWork, error) {
	dayOfWorkEntity := entities.DayOfWork{
		Day:       dayOfWork.Day,
		StartTime: dayOfWork.StartTime,
		EndTime:   dayOfWork.EndTime,
		IsWorkDay: dayOfWork.IsWorkDay,
	}
	result, err := s.repository.Update(id, dayOfWorkEntity)
	if err != nil {
		return nil, err
	}
	return &models.DayOfWork{
		ID:        result.ID,
		Day:       result.Day,
		StartTime: result.StartTime,
		EndTime:   result.EndTime,
		IsWorkDay: result.IsWorkDay,
	}, nil
}

func (s dayOfWorkService) Delete(id uint) error {
	return s.repository.Delete(id)
}
